VERSION=development

default:
	@echo "=============Building binaries============="

	# Linux 386
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/linux_386/simple-proxy main.go
	cp LICENSE dist/linux_386/LICENSE

	# Linux amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/linux_amd64/simple-proxy main.go
	cp LICENSE dist/linux_amd64/LICENSE

	# Linux arm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/linux_arm/simple-proxy main.go
	cp LICENSE dist/linux_arm/LICENSE

	# Linux arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/linux_arm64/simple-proxy main.go
	cp LICENSE dist/linux_arm64/LICENSE

	# Darwin amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/darwin_amd64/simple-proxy main.go
	cp LICENSE dist/darwin_amd64/LICENSE

	# Darwin arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'main.Version=$(VERSION)' -s -w" -o dist/darwin_arm64/simple-proxy main.go
	cp LICENSE dist/darwin_arm64/LICENSE

	# Windows 386
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-X 'main.Version=$(VERSION)'" -o dist/windows_386/simple-proxy.exe main.go
	cp LICENSE dist/windows_386/LICENSE

	# Windows amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=$(VERSION)'" -o dist/windows_amd64/simple-proxy.exe main.go
	cp LICENSE dist/windows_amd64/LICENSE

zip:
	@echo "=============Zipping binaries============="
	zip -r -j dist/simple-proxy_linux_386.zip dist/linux_386
	zip -r -j dist/simple-proxy_linux_amd64.zip dist/linux_amd64
	zip -r -j dist/simple-proxy_linux_arm.zip dist/linux_arm
	zip -r -j dist/simple-proxy_linux_arm64.zip dist/linux_arm64
	zip -r -j dist/simple-proxy_darwin_amd64.zip dist/darwin_amd64
	zip -r -j dist/simple-proxy_darwin_arm64.zip dist/darwin_arm64
	zip -r -j dist/simple-proxy_windows_386.zip dist/windows_386
	zip -r -j dist/simple-proxy_windows_amd64.zip dist/windows_amd64

lint:
	@echo "=============Linting============="
	staticcheck ./...

format:
	@echo "=============Formatting============="
	gofmt -s -w .
	go mod tidy

test:
	@echo "=============Running unit tests============="
	go test ./...
