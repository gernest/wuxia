
test:generate
	@go test  ./wuxia
	@go test  ./db
	@go test  ./migration
	@go test  ./models
	@go test  ./session
	@go test  ./api

setup:
	@go get -v github.com/jteeuwen/go-bindata/...

migration/data.gen.go:$(shell find migration/scripts -type f)
	@echo "Generating migration scripts bindata"
	@go generate ./migration

themes/data.gen.go:$(shell find themes/theme -type f)
	@echo "Generating themes bindata"
	@go generate ./themes

wuxia/data.gen.go:$(shell find wuxia/js -type f)
	@echo "Generating generator bindata"
	@go generate ./wuxia

generate: migration/data.gen.go themes/data.gen.go wuxia/data.gen.go
	@echo "Done generate bindata"

cover:
	./coverage.sh
