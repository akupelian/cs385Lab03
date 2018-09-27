.PHONY: gopp

goapp: 
	docker build -t goapp .
	docker run -d --name goapp goapp

clean:
	docker stop goapp
	docker rm goapp

prune:
	docker system prune -a
