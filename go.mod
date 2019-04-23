module github.com/victoryang/api-gateway

go 1.12

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894

require (
	github.com/gorilla/mux v1.7.1
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/urfave/negroni v1.0.0
	gopkg.in/yaml.v2 v2.2.2
)
