module github.com/gopaytech/go-commons

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/containerd/containerd v1.3.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.3.0
	k8s.io/client-go v0.0.0-20191111061043-a56922badea0
)

//https://github.com/docker/engine/releases/tag/v19.03.5
replace github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
