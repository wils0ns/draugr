air:
	@air

run: build
	@./bin/draugr

install:
	@go get ./...
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod tidy
	@go mod download
	@go install github.com/cosmtrek/air@latest

	@npm install -D tailwindcss
	@npm i -D daisyui@latest

# day1: install

build:
	@npx tailwindcss -i view/css/app.css -o public/app.css
	@templ generate view
	@go build -o bin/draugr
