module github.com/gopaytech/go-commons

go 1.16

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
	github.com/fatih/color v1.10.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/goreleaser/nfpm v1.10.3
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/sirupsen/logrus v1.7.0
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	gopkg.in/alexcesaro/statsd.v2 v2.0.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
	k8s.io/client-go v0.20.0
)

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200916142827-bd33bbf0497b+incompatible
