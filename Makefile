.PHONY: build serve clean install

# Build the site
build:
	go run ./cmd/build

# Build and serve locally
serve:
	go run ./cmd/build -serve -port 8080

# Clean output directory
clean:
	rm -rf output

# Install dependencies
install:
	go mod tidy

# Build the generator binary
binary:
	go build -o build ./cmd/build

# Deploy to GitHub Pages (example)
deploy: build
	@echo "Output ready in ./output"
	@echo "Deploy with: cd output && git init && git add . && git commit -m 'deploy' && git push -f git@github.com:USER/USER.github.io.git main"
