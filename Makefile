run: build
	@templ generate view
	@./bin/sentryeye

day1:
	# Install nodejs and npm

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod tidy
	@go mod download

build:
	@go build -o bin/sentryeye
