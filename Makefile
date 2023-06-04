.DEFAULT_GOAL := run

build:
		docker build -f back-end/Dockerfile -t url_shortener:local ./back-end
.PHONY:build

run: build
		docker compose up 
.PHONY:run