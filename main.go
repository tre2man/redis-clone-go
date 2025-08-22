package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type opType int

const (
	opGet opType = iota
	opSet
	opDel
)

type cmd struct {
	op    opType
	key   string
	value string
	resp  chan string // 응답 전달용
}

type store struct {
	m map[string]string
}

func storageLoop(in <-chan cmd) {
	s := store{m: make(map[string]string)}
	for c := range in {
		switch c.op {
		case opGet:
			c.resp <- s.m[c.key]
		case opSet:
			s.m[c.key] = c.value
			c.resp <- "OK"
		case opDel:
			if _, exists := s.m[c.key]; exists {
				delete(s.m, c.key)
				c.resp <- "OK"
			} else {
				c.resp <- "ERR key not found"
			}
		}
	}
}

func handleConn(c net.Conn, dispatch chan<- cmd) {
	defer c.Close()
	r := bufio.NewScanner(c)
	// 타임아웃/쓰기버퍼/에러처리는 생략(프로덕션에선 꼭 추가)
	for r.Scan() {
		line := strings.TrimSpace(r.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 3)
		op := strings.ToUpper(parts[0])
		respCh := make(chan string, 1)

		switch op {
		case "GET":
			if len(parts) < 2 {
				fmt.Fprintln(c, "ERR wrong args")
				continue
			}
			dispatch <- cmd{op: opGet, key: parts[1], resp: respCh}
			fmt.Fprintln(c, <-respCh)
		case "SET":
			if len(parts) < 3 {
				fmt.Fprintln(c, "ERR wrong args")
				continue
			}
			dispatch <- cmd{op: opSet, key: parts[1], value: parts[2], resp: respCh}
			fmt.Fprintln(c, <-respCh)
		case "DEL":
			if len(parts) < 2 {
				fmt.Fprintln(c, "ERR wrong args")
				continue
			}
			dispatch <- cmd{op: opDel, key: parts[1], resp: respCh}
			fmt.Fprintln(c, <-respCh)
		case "QUIT":
			fmt.Fprintln(c, "OK")
			return
		case "PING":
			fmt.Fprintln(c, "PONG")
		default:
			fmt.Fprintln(c, "ERR unknown cmd")
		}
	}
}

func main() {
	dispatch := make(chan cmd, 1024) // 백프레셔
	go storageLoop(dispatch)

	ln, err := net.Listen("tcp", ":6380")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening :6380")
	var wg sync.WaitGroup
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			handleConn(conn, dispatch)
		}()
	}
}
