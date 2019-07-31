module github.com/magneticio/istio-client-go

go 1.12

//replace istio.io/api => github.com/rcernich/istio-api v0.0.0-20190211150719-23148abe2cc6

require (
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.2.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/rcernich/istio-api v0.0.0-20190115020421-40a08a31eaf1 // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/stretchr/testify v1.3.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/grpc v1.22.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	istio.io/api v0.0.0-20190730193711-5ad2b2c986fc
	k8s.io/api v0.0.0-20190722141453-b90922c02518 // indirect
	k8s.io/apimachinery v0.0.0-20190719140911-bfcf53abc9f8
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190712204705-3dccf664f023 // indirect
)
