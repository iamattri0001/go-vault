package injector

import "go-vault/service"

type DependencyInjector struct {
	Service *service.Service
}

var SingletonInjector *DependencyInjector

func NewDependencyInjector(options ...func(di *DependencyInjector)) *DependencyInjector {
	injector := &DependencyInjector{}
	for _, option := range options {
		option(injector)
	}
	SingletonInjector = injector
	return injector
}

func GetSingletonInjector() *DependencyInjector {
	if SingletonInjector == nil {
		panic("DependencyInjector is not initialized")
	}
	return SingletonInjector
}
