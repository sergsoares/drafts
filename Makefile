.PHONY: help new

# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

# List and preview descriptions for each project
show: 
	find * -name description.md | fzf --preview 'cat {}'

# Create a new project
new:
	@read -p "What is the project name: " project; \
	mkdir $$project && cd $$project && touch Makefile && touch description.md
