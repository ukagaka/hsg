package model

import "hxsg/internal/model/User"

var M = &ModuleManager{}

type ModuleManager struct {
	//这里初始化各种注册各种 model

	UserModel User.UserModel
	DefaultModuleManager
}

func (this *ModuleManager) RegisterBaseModel() {

	//这里注册各种 model

	err := this.DefaultModuleManager.Init()
	if err != nil {
		panic(err)
	}
}

func (this *ModuleManager) AppendModule(module Module) Module {
	this.Modules = append(this.Modules, module)
	return module
}
