CC=rustc -A warnings

default: build

build: main.rs
	${CC} $<

run: build
	./main input

test: build
	./main input.test
