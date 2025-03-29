all: build

build:
	@echo "Building the project..."
	@mkdir -p build
	@go build -o build/tt-cli
	@echo "Build complete."

install:
	@echo "Installing the project..."
	@mkdir -p $(GOPATH)/bin
	@cp build/tt-cli $(GOPATH)/bin/
	@echo "Installation complete."
	@echo "You can run the tool using 'tt-cli' command."
	@echo "To uninstall, run 'make uninstall'."
uninstall:
	@echo "Uninstalling the project..."
	@rm -f $(GOPATH)/bin/tt-cli
	@echo "Uninstallation complete."

clean:
	@echo "Cleaning up..."
	@rm -rf build
	@echo "Done."