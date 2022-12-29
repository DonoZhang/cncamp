export tag=v1.0
export image=dono1992/cnhttpserver

release:
	echo "building httpserver container"
	docker build --no-cache -t ${image}:${tag} --build-arg version=${tag} .

publish: release
	echo "pushing ${image}:${tag}"
	docker push ${image}:${tag}