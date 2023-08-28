package svc

import (
	"seal/server/internal/config"

	shell "github.com/ipfs/go-ipfs-api"
)

type ServiceContext struct {
	Config config.Config
	Sh     *shell.Shell
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sh:     shell.NewShell(c.Ipfs),
	}
}
