package plugins

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

type Service interface {
	csi.IdentityServer
	csi.ControllerServer
	csi.NodeServer
}

// Factory method
func New(Opts interface{}) (*Service, error) {
	// TODO
}
