.PHONY: help new

# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

ls:
	ls -tl | fzf

# List and preview descriptions for each project
show: 
	find * -name description.md | fzf --preview 'cat {}'

# Create a new project (formatting spaces to hifen)
new:
	@read -p "What is the project name: " input; \
	project=`echo $$input | sed -e 's/\s$$//g' | tr ' ' '-' | tr '[:upper:]' '[:lower:]' ` ; \
	mkdir $$project && cd $$project && touch Makefile && touch description.md; \
	echo "$$project created, enter with $ cd $$project"