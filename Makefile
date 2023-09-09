new:
	@read -p "What is the project name: " project; \
	mkdir $$project && cd $$project && touch Makefile
