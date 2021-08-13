package deb

import (
	"fmt"

	"github.com/goreleaser/nfpm"
)

type Config struct {
	Name        string
	Version     string
	Destination string
	PostRemove  string
	Arch        string
	Source      string
}

func (c Config) ConvertToNFPMConfig() (config *nfpm.Info) {
	config = &nfpm.Info{
		Name:        c.Name,
		Arch:        c.Arch,
		Version:     c.Version,
		Bindir:      c.Destination,
		Maintainer:  "Gopay-Systems",
		Description: fmt.Sprintf("%s_%s", c.Name, c.Version),
		Overridables: nfpm.Overridables{
			Scripts: nfpm.Scripts{
				PostRemove: c.PostRemove,
			},
		},
	}
	return
}
