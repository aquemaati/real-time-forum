# Variables
BINARY_NAME=myapp
MODULE_PATH=backend
MAIN_FILE=main.go

# Targets
all: run

build:
	cd $(MODULE_PATH) && go build -o ../$(BINARY_NAME) $(MAIN_FILE)

run: build
	@trap '$(MAKE) clean' EXIT INT TERM; \
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run clean
