package container

import (
	"reflect"
)

type Container struct {
	services map[reflect.Type]interface{}
}

func NewContainer() *Container {
	return &Container{
		services: make(map[reflect.Type]interface{}),
	}
}

func (c *Container) Register(service interface{}) {
	c.services[reflect.TypeOf(service)] = service
}

func (c *Container) Resolve(service interface{}) interface{} {
	return c.services[reflect.TypeOf(service)]
}
