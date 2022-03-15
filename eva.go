package main

import "regexp"

// ------------ inplementation -------------

type Eva struct {
	envs Enviroment
}

func Constructor(global Enviroment) Eva {
	return Eva{
		envs: global,
	}
}

func (eva Eva) Evaluate(exp interface{}) interface{} {
	return eval(exp, eva.envs)
}

func eval(exp interface{}, env Enviroment) interface{} {
	if isVariableName(exp) {
		return env.lookup(exp.(string))
	}
	// self evaluating expressions:
	if isNumber(exp) {
		return exp.(int)
	}
	if isString(exp) {
		return exp.(string)[1 : len(exp.(string))-1]
	}
	if multiArgs(exp) {
		if exp.([]interface{})[0] == "+" {
			return eval(exp.([]interface{})[1], env).(int) + eval(exp.([]interface{})[2], env).(int)
		}

		if exp.([]interface{})[0] == "-" {
			return eval(exp.([]interface{})[1], env).(int) - eval(exp.([]interface{})[2], env).(int)
		}

		if exp.([]interface{})[0] == "*" {
			return eval(exp.([]interface{})[1], env).(int) * eval(exp.([]interface{})[2], env).(int)
		}

		if exp.([]interface{})[0] == "/" {
			return eval(exp.([]interface{})[1], env).(int) / eval(exp.([]interface{})[2], env).(int)
		}

		// variable declarations
		if exp.([]interface{})[0] == "var" {
			return env.define(exp.([]interface{})[1].(string), eval(exp.([]interface{})[2], env))
		}
	}

	panic("not implemented")
}

func isNumber(exp interface{}) bool {
	switch exp.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

func isString(exp interface{}) bool {
	switch exp.(type) {
	case string:
		return exp.(string)[0] == '"'
	}
	return false
}

func multiArgs(exp interface{}) bool {
	switch exp.(type) {
	case []interface{}:
		return true
	}
	return false
}

func isVariableName(exp interface{}) bool {
	switch exp.(type) {
	case string:
		match, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]*$", exp.(string))
		return match
	}
	return false
}
