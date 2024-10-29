build:
	docker build -t temperature-system .

run:
	docker run --rm -p 8080:8080 temperature-system
