program_name = proxima
dist_dir = dist/

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(dist_dir)$(program_name)-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(dist_dir)$(program_name)-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(dist_dir)$(program_name)-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(dist_dir)$(program_name)-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(dist_dir)$(program_name)-darwin-arm64 main.go

clean:
	@rm -rf $(dist_dir)*

test:
	@sudo mv $(dist_dir)$(program_name)-darwin-arm64 /usr/local/bin/$(program_name)
	@chmod +x /usr/local/bin/$(program_name)
	@$(program_name) version

.PHONY: compile clean test
