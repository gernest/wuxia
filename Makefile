
test:generate
	@go test -v ./wuxia
	@go test -v ./api

setup:
	@go get -v github.com/jteeuwen/go-bindata/...

migration/data.go:$(shell find migration/scripts -type f)
	@echo "Generating migration scripts bindata"
	@go generate ./migration

generate: migration/data.go
