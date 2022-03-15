package main

import "fmt"

type Enviroment struct {
	vals map[string]interface{}
}

func (env Enviroment) define(name string, value interface{}) interface{} {
	env.vals[name] = value
	return value
}

func (env Enviroment) lookup(name string) interface{} {
	if v, ok := env.vals[name]; ok {
		return v
	}
	panic(fmt.Sprintf("%s is not defined.", name))
}
