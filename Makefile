protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/layer5io/meshery/master/meshes/meshops.proto

proto:	
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

docker:
	docker build -t layer5/meshery-maesh .

docker-run:
	(docker rm -f meshery-maesh) || true
	docker run --name meshery-maesh -d \
	-p 10000:10000 \
	-e DEBUG=true \
	layer5/meshery-maesh

run:
	DEBUG=true go run main.go