# errgroup 丢错

internal/parallel/run.go 忽略 g.Wait 错误

## 环境

- 镜像：benzhi.Dockerfile 基于 golang:1.22（官方多架构）
- go.mod 语言版本：go 1.22
- 容器内使用镜像自带工具链即可

## 标准命令

```bash
go build ./...
go test ./... -count=1
go vet ./...
```

## 构建评测镜像（须双架构）

验证请用 `bash -c`（勿用 `bash -lc`）。

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh go-errgroup-wait linux/amd64
docker run --platform linux/amd64 --rm go-errgroup-wait:latest bash -c 'go build ./... && go test ./... -count=1'

./build_benzhi_docker.sh go-errgroup-wait linux/arm64
docker run --platform linux/arm64 --rm go-errgroup-wait:latest bash -c 'go build ./... && go test ./... -count=1'
```

构建阶段已处理依赖下载；容器内不应再出现 downloading ...。
