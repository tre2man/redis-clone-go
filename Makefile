# RedisCloneGo Makefile

# 변수 정의
BINARY_NAME=redisclonego
BUILD_DIR=build
MAIN_FILE=main.go
GO=go
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# 기본 타겟
.PHONY: all build run clean help

# 기본 타겟 (make만 실행했을 때)
all: build

# 빌드 타겟 - README.md에 따라 redisclonego 바이너리 생성
build:
	@echo "Building RedisCloneGo..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# 실행 타겟
run: build
	@echo "Running RedisCloneGo..."
	./$(BUILD_DIR)/$(BINARY_NAME)



# 테스트 커버리지 확인
test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# 정리 (빌드 파일 및 바이너리 제거)
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out
	@echo "Clean completed"

# 의존성 다운로드
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download

# 의존성 정리
deps-clean:
	@echo "Cleaning dependencies..."
	$(GO) mod tidy

# 린트 실행
lint:
	@echo "Running linter..."
	$(GO) vet ./...
	golangci-lint run

# 포맷팅
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# 도움말
help:
	@echo "Available targets:"
	@echo "  build        - Build the binary (default)"
	@echo "  run          - Build and run the binary"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download dependencies"
	@echo "  deps-clean   - Clean dependencies"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  help         - Show this help message"
