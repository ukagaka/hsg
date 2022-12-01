package model

import (
	"fmt"
	"hxsg/logger"
	"strings"
	"sync"
)

type Module interface {
	Init() error
	Start() error
	Run()
	Stop()
}

type DefaultModuleManager struct {
	Module
	Modules []Module
}

func (this *DefaultModuleManager) Init() error {
	for i := 0; i < len(this.Modules); i++ {
		clsName := fmt.Sprintf("%T", this.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		logger.Info(clsName[dotIndex:] + " Init")
		err := this.Modules[i].Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *DefaultModuleManager) Start() error {
	for i := 0; i < len(this.Modules); i++ {
		err := this.Modules[i].Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *DefaultModuleManager) Run() {
	for i := 0; i < len(this.Modules); i++ {
		clsName := fmt.Sprintf("%T", this.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		logger.Info(clsName[dotIndex:] + " Run")
		this.Modules[i].Run()
	}
}

func (this *DefaultModuleManager) Stop() {
	var wg sync.WaitGroup
	for i := 0; i < len(this.Modules); i++ {
		clsName := fmt.Sprintf("%T", this.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		logger.Info(clsName[dotIndex:] + " Stop")
		wg.Add(1)
		go func(module Module) {
			module.Stop()
			wg.Done()
		}(this.Modules[i])
	}
	wg.Wait()
}
