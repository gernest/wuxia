
test:generate
	@go test  ./wuxia
	@go test  ./db
	@go test  ./migration
	@go test  ./models
	@go test  ./session
	@go test  ./base

setup:
	@go get -v github.com/jteeuwen/go-bindata/...

migration/data.gen.go:$(shell find migration/scripts -type f)
	@echo "Generating migration scripts bindata"
	@go generate ./migration


wuxia/data.gen.go:$(shell find wuxia/js -type f)
	@echo "Generating generator bindata"
	@go generate ./wuxia

data/data.gen.go:$(shell find public/  -type f)
	@echo "generating data for public files"
	@go-bindata -o data/data.gen.go\
		-pkg data -prefix public/ public/...

views/data.gen.go:$(shell find templates/  -type f)
	@echo "generating data for templates"
	@go-bindata -o views/data.gen.go\
		-pkg views -prefix templates/ templates/...
generate: migration/data.gen.go  wuxia/data.gen.go views/data.gen.go data/data.gen.go

	@echo "Done generate bindata"

cover:
	./coverage.sh

build:
	@go build -o wu ./cmd/wuxia
