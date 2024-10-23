# 设置 shell 为 PowerShell
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

# 定义 Go 命令
GO := go

# 定义删除命令
RM := if exist
rpc-service:
	$(GO) build $(LDFLAGS)
.PHONY: rpc-service

clean:
	$(RM) rpc-service

test:
	$(GO) test -v ./...

lint:
	golangci-lint run ./...