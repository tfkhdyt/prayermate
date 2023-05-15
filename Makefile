# Go program name
PROGRAM_NAME=prayermate

# Go compiler command
GO=go

# Install directory
INSTALL_DIR=/usr/local/bin

# Share directory
SHARE_DIR=/usr/local/share/prayermate

# Assets directory
ASSETS_DIR=assets

# Completion script paths
BASH_COMPLETION=/etc/bash_completion.d/$(PROGRAM_NAME)
ZSH_COMPLETION=/usr/share/zsh/site-functions/_$(PROGRAM_NAME)
FISH_COMPLETION=/usr/share/fish/vendor_completions.d/$(PROGRAM_NAME).fish

all: build

build:
	$(GO) build -buildvcs=false -o $(PROGRAM_NAME) .

install: build
	install -m 0755 $(PROGRAM_NAME) $(INSTALL_DIR)
	mkdir -p $(SHARE_DIR)
	cp -R $(ASSETS_DIR) $(SHARE_DIR)
	$(GO) run . completion bash > $(BASH_COMPLETION)
	$(GO) run . completion zsh > $(ZSH_COMPLETION)
	$(GO) run . completion fish > $(FISH_COMPLETION)

uninstall:
	rm -f $(INSTALL_DIR)/$(PROGRAM_NAME)
	rm -rf $(SHARE_DIR)
	rm -f $(BASH_COMPLETION)
	rm -f $(ZSH_COMPLETION)
	rm -f $(FISH_COMPLETION)

clean:
	rm -f $(PROGRAM_NAME)

.PHONY: all build install clean
