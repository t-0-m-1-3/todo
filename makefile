all: dist

dist:
	make -C frontend
	go build -o dist/todo

clean:
	@rm -rf dist
	@make -C frontend clean