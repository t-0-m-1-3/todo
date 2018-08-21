all: dist

dist:
	make -C frontend
	packr build -o dist/todo

start: dist
	dist/todo

clean:
	@rm -rf dist
	@make -C frontend clean