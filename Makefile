
test:generate
	@go test -v ./wuxia
	@go test -v ./api

setup:
	@go get -v github.com/jteeuwen/go-bindata/...

migration/data.go:$(shell find migration/scripts -type f)
	@echo "Generating migration scripts bindata"
	@go generate ./migration

themes/data.go:$(shell find themes/theme -type f)
	@echo "Generating themes bindata"
	@go generate ./themes

generate: migration/data.go themes/data.go
