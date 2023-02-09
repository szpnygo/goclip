version := $(shell /bin/date "+%Y-%m-%d %H:%M")

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goclip main.go
	$(if $(shell command -v upx), upx goclip)
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goclip main.go
	$(if $(shell command -v upx), upx goclip)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goclip.exe main.go
	$(if $(shell command -v upx), upx goclip.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goclip main.go
	$(if $(shell command -v upx), upx goclip)
