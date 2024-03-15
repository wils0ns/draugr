run: build
	@./bin/sentryeye

day1:
	@npm install -D tailwindcss
	@npm i -D daisyui@latest

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod tidy
	@go mod download

build:
	@npx tailwindcss -i view/css/app.css -o public/app.css
	@templ generate view
	@go build -o bin/sentryeye
