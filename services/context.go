package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type ServiceMap map[reflect.Type]reflect.Value

func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(ServiceMap))
	}
	return c
}
