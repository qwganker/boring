go mod tidy
go build  -ldflags "-X main.GinMode=release" -o boring-agent ./cmd/agent/
go build  -ldflags "-X main.GinMode=release" -o boring-server ./cmd/server/
go build  -ldflags "-X main.GinMode=release" -o boring-jobworker ./cmd/jobworker/
