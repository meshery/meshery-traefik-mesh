protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/layer5io/meshery/master/meshes/meshops.proto

proto:	
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

docker:
	docker build -t layer5/meshery-traefik-mesh .

docker-run:
	(docker rm -f meshery-traefik-mesh) || true
	docker run --name meshery-traefik-mesh -d \
	-p 10000:10000 \
	-e DEBUG=true \
	layer5/meshery-traefik-mesh

run:
	DEBUG=true go run main.go


error:
	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers
