APP_NAME := jit
DIST_DIR := dist
GOOS_LIST := linux darwin windows
GOARCH_LIST := amd64 arm64

.PHONY: all clean build

all: clean build

clean:
	@echo "🧹 Cleaning build output..."
	rm -rf $(DIST_DIR)

build:
	@echo "📦 Building binaries..."
	@mkdir -p $(DIST_DIR)
	@for goos in $(GOOS_LIST); do \
	  for goarch in $(GOARCH_LIST); do \
	    ext=""; \
	    if [ "$$goos" = "windows" ]; then ext=".exe"; fi; \
	    output="$(DIST_DIR)/$(APP_NAME)-$$goos-$$goarch$$ext"; \
	    echo "🚧 $$goos/$$goarch -> $$output"; \
	    GOOS=$$goos GOARCH=$$goarch go build -o $$output main.go; \
	  done \
	done
