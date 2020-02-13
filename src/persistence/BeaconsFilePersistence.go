package persistence

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cmperist "github.com/pip-services3-go/pip-services3-data-go/persistence"
)

type BeaconsFilePersistence struct {
	BeaconsMemoryPersistence
	persister *cmperist.JsonFilePersister
}

func NewBeaconsFilePersistence(path string) *BeaconsFilePersistence {
	bfp := BeaconsFilePersistence{}
	bfp.BeaconsMemoryPersistence = *NewBeaconsMemoryPersistence()
	bfp.persister = cmperist.NewJsonFilePersister(bfp.Prototype, path)
	bfp.Loader = bfp.persister
	bfp.Saver = bfp.persister
	return &bfp
}

func (c *BeaconsFilePersistence) Configure(config *cconf.ConfigParams) {
	c.BeaconsMemoryPersistence.Configure(config)
	c.persister.Configure(config)
}
