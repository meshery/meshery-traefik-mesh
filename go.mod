module github.com/layer5io/meshery-traefik-mesh

go 1.15

replace github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200806180306-b7e46afd657f

require (
	github.com/layer5io/meshery-adapter-library v0.1.18
	github.com/layer5io/meshkit v0.2.14
	github.com/layer5io/service-mesh-performance v0.3.3
	gopkg.in/yaml.v2 v2.4.0
)
