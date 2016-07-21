
test:generate
	@go test  ./wuxia
	@go test  ./db
	@go test  ./migration
	@go test  ./models
	@go test  ./api

setup:
	@go get -v github.com/jteeuwen/go-bindata/...

migration/data.go:$(shell find migration/scripts -type f)
	@echo "Generating migration scripts bindata"
	@go generate ./migration

themes/data.go:$(shell find themes/theme -type f)
	@echo "Generating themes bindata"
	@go generate ./themes

generate: migration/data.go themes/data.go
	@echo "Done generate bindata"
