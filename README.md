# RedisCloneGo

## 소개

redis 프로젝트를 go 로 구현한 프로젝트입니다. 최종적으로는 redis 의 테스트 코드를 해당 프로젝트에서 모두 통과하게 하는 것이 목표힙니다.

## 구현된 기능

```
SET key value
GET key
DEL key
```

## 실행 방법

### 기본 빌드 및 실행
```shell
go build -o redisgo build/main.go
./redisgo
```

### Makefile 사용 (권장)
프로젝트에는 편리한 Makefile이 포함되어 있습니다.

#### 기본 사용법
```shell
# 바이너리 빌드
make build

# 빌드 후 실행
make run

# 빌드 파일 정리
make clean
```

#### 사용 가능한 모든 명령어
```shell
make help
```

#### 주요 타겟
- **`build`**: `build/redisgo` 바이너리 파일 생성 (기본 타겟)
- **`run`**: 빌드 후 바이너리 실행
- **`clean`**: 빌드 파일 및 바이너리 제거
- **`test-coverage`**: 테스트 커버리지 확인
- **`deps`**: 의존성 다운로드
- **`deps-clean`**: 의존성 정리
- **`lint`**: 코드 린트 실행
- **`fmt`**: 코드 포맷팅
- **`help`**: 사용 가능한 모든 명령어 표시