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
	$(GO) build -buildvcs=false -ldflags="-w -s" -o $(PROGRAM_NAME) .
	./$(PROGRAM_NAME) completion bash > $(PROGRAM_NAME).bash
	./$(PROGRAM_NAME) completion zsh > _$(PROGRAM_NAME).zsh
	./$(PROGRAM_NAME) completion fish > $(PROGRAM_NAME).fish

install: build
	mkdir -p $(SHARE_DIR) $(INSTALL_DIR)
	install -Dm755 $(PROGRAM_NAME) $(INSTALL_DIR)
	install -Dm644 $(PROGRAM_NAME).bash $(BASH_COMPLETION)
	install -Dm644 _$(PROGRAM_NAME).zsh $(ZSH_COMPLETION)
	install -Dm644 $(PROGRAM_NAME).fish $(FISH_COMPLETION)
	cp -R $(ASSETS_DIR) $(SHARE_DIR)

uninstall:
	rm -f $(INSTALL_DIR)/$(PROGRAM_NAME)
	rm -rf $(SHARE_DIR)
	rm -f $(BASH_COMPLETION)
	rm -f $(ZSH_COMPLETION)
	rm -f $(FISH_COMPLETION)

clean:
	rm -f $(PROGRAM_NAME) $(PROGRAM_NAME).bash _$(PROGRAM_NAME).zsh $(PROGRAM_NAME).fish

.PHONY: all build install uninstall clean
