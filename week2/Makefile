export tag=v1.0

build:
	echo "building httpserver binary"
	GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t dono1992/cnhttpserver:${tag} --build-arg version=${tag} .

push: release
	echo "pushing dono1992/cnhttpserver:${tag}"
	docker push dono1992/cnhttpserver:${tag}