ifeq "$(origin BUILD_FLAG)" "undefined"
	BUILD_FLAG=-debug
endif



test:generate
	@go test  ./gen
	@go test  ./db
	@go test  ./migration
	@go test  ./models
	@go test  ./session
	@go test  ./base

setup:
	@go get -v github.com/jteeuwen/go-bindata/...
	@go get github.com/Unknwon/bra

migration/data.gen.go:$(shell find migration/scripts -type f)
	@echo "generating: migration/data.gen.go "
	@go-bindata  $(BUILD_FLAG) \
		-o migration/data.gen.go \
		-pkg migration  \
		-prefix migration/ \
		migration/scripts/...


gen/data.gen.go:$(shell find gen/js -type f)
	@echo "generating: geb/data.gen.go "
	@go-bindata $(BUILD_FLAG) \
		-o gen/data.gen.go \
		-pkg gen \
		-prefix gen/ \
		gen/js/...

data/data.gen.go:$(shell find public/  -type f)
	@echo "generating: data/data.gen.go "
	@go-bindata $(BUILD_FLAG) \
		-o data/data.gen.go\
		-pkg data -prefix public/ public/...

views/data.gen.go:$(shell find templates/  -type f)
	@go-bindata $(BUILD_FLAG) \
		-o views/data.gen.go\
		-pkg views -prefix templates/ templates/...

generate: migration/data.gen.go  gen/data.gen.go views/data.gen.go data/data.gen.go
	@echo "Done generate bindata"

cover:
	./coverage.sh

build:generate
	@go build -o wu ./cmd/wuxia
