SRC_DIR := cmd

# Define the output folder for binary files
BIN_DIR := bin

# Find all subdirectories in the SRC_DIR
SUBDIRS := $(wildcard $(SRC_DIR)/*)

# Generate the binary file names based on the subdirectory names
BINARIES := $(patsubst $(SRC_DIR)/%,$(BIN_DIR)/%,$(SUBDIRS))

# The default target builds all binaries
all: copy_output $(BINARIES)

# Define a pattern rule to build each binary from its corresponding source directory
$(BIN_DIR)/%: $(SRC_DIR)/%
	@mkdir -p $(BIN_DIR)
	go build -o $@ $</main.go

copy_output:
	@mkdir -p $(BIN_DIR)
	cp ./app-conf.yml $(BIN_DIR)/app-conf.yml
	cp ./app-conf-dev.yml $(BIN_DIR)/app-conf-dev.yml
	cp -rf ./migrations $(BIN_DIR)/migrations

	cp -rf ./public $(BIN_DIR)/public
	cp -rf ./templates $(BIN_DIR)/templates

clean:
	rm -rf $(BIN_DIR)

migrations-setup: all
	./bin/migrate setup
	
migrations-up: all
	./bin/migrate up
	
migrations-down: all
	./bin/migrate down

.PHONY: all clean


