.PHONY: test build run


test:
	make -C person-post test

build: 
	make -C person-post build

run:
	make -C person-post run
