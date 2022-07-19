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

## Build and run Adapter locally
run:
	go$(v) mod tidy -compat=1.17; \
	DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go

run-force-dynamic-reg:
	FORCE_DYNAMIC_REG=true DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go

error:
	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers
