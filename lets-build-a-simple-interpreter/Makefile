.PHONY: help new

# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

# Execute calc1
calc1:
	@python3 calc1.py

# Execute calc2
calc2:
	@python3 calc2.py

# Execute calc3
calc3:
	@python3 calc3.py

calc3-copy:
	@python3 'calc3 copy.py'
	