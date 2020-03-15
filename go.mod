module github.com/talos-systems/machined

go 1.13

require (
	github.com/containerd/containerd v1.3.2
	github.com/containerd/cri v1.11.1
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/kubernetes-incubator/bootkube v0.14.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.1
	github.com/stretchr/testify v1.5.0
	github.com/talos-systems/go-procfs v0.0.0-20200219015357-57c7311fdd45
	github.com/talos-systems/grpc-proxy v0.2.0
	github.com/talos-systems/talos v0.3.2
	golang.org/x/sys v0.0.0-20200107162124-548cf772de50
	google.golang.org/grpc v1.26.0
	k8s.io/apiextensions-apiserver v0.17.4 // indirect
	k8s.io/cri-api v0.0.0-20191121183020-775aa3c1cf73
	k8s.io/kubelet v0.17.4 // indirect
)
