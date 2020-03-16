image=/timerbot:latest
remote=

build:
	go build
	docker build -t ${image} .
	docker push ${image}

install:
	make build
	ssh ${remote} docker pull ${image}
	ssh ${remote} docker stop timerbot || true
	ssh ${remote} docker rm timerbot || true
	ssh ${remote} docker run -d --restart=always --name=timerbot ${image}
