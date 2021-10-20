OUT = ./build
SRC = ./cmd/app/main.go
EXEC = app
COVER = cover.out
COVER_HTML = cover.html

.DEFAULT_GOAL: build

.PHONY: build test clean html-cover
build:
	mkdir -p $(OUT)
	go build --o $(OUT)/$(EXEC) $(SRC)

test: build
	go test -v -cover ./... -coverprofile=$(OUT)/$(COVER)

html-cover: test
	go tool cover -html=$(OUT)/$(COVER) -o $(OUT)/$(COVER_HTML)
clean:
	rm -rf $(OUT)