.PHONY: help new

# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

# List directories with folder names
ls:
	ls -d -t */ | fzf --preview-window=up --preview 'stat -f "%Sm %N" {}*' 
	
# Create a new project (formatting spaces to hifen)
create:
	@read -p "What is the project name: " input; \
	project=`echo $$input | sed -e 's/\s$$//g' | tr ' ' '-' | tr '[:upper:]' '[:lower:]' ` ; \
	mkdir $$project && cd $$project && touch Makefile && touch description.md; \
	echo "$$project created, enter with $ cd $$project"