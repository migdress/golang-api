.PHONY: test build run


test:
	make -C person-post test

build: 
	make -C person-post build

run:
	make -C person-post run

dockerrun:
	docker build -t person-post:v1 .
	docker run -it --network="host" person-post:v1