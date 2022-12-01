package base

import (
	"hxsg/config"
	"hxsg/internal/model"
	"hxsg/logger"
	"hxsg/orm"
)

type Container struct {
}

var container = NewContainerManager()

func GetContainer() *Container {
	return container
}

func NewContainerManager() *Container {
	return &Container{}
}

func (this *Container) InitData() {

	config.Init()
	logger.Init()
	orm.Init()
	model.M.RegisterBaseModel()
}
