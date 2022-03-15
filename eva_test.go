package main

import "testing"

func TestEval(t *testing.T) {
	t.Log("start test")

	var global Enviroment = Enviroment{vals: map[string]interface{}{
		"null":    nil,
		"true":    true,
		"false":   false,
		"VERSION": "0.0.1",
	}}

	eva := Constructor(global)
	if eva.Evaluate(1) != 1 {
		t.Error("eval number fail.")
	}

	if eva.Evaluate("\"hello\"") != "hello" {
		t.Error("eval string fail.")
	}

	if eva.Evaluate([]interface{}{"+", 1, 5}) != 6 {
		t.Error("add string fail.")
	}

	if eva.Evaluate([]interface{}{"+", 1, []interface{}{"+", 2, 3}}) != 6 {
		t.Error("add expressions string fail.")
	}

	if eva.Evaluate([]interface{}{"+", 1, []interface{}{"-", 2, 3}}) != 0 {
		t.Error("sub expressions string fail.")
	}

	if eva.Evaluate([]interface{}{"*", 2, 5}) != 10 {
		t.Error("multi string fail.")
	}

	if eva.Evaluate([]interface{}{"/", 10, 5}) != 2 {
		t.Error("devide string fail.")
	}

	if eva.Evaluate([]interface{}{"var", "foo", 5}) != 5 {
		t.Error("define variable fail.")
	}

	if eva.Evaluate([]interface{}{"var", "foo", []interface{}{"+", 2, 3}}) != 5 {
		t.Error("define variable fail.")
	}

	if eva.Evaluate("foo") != 5 {
		t.Error("lookup variable fail.")
	}

	if eva.Evaluate("VERSION") != "0.0.1" {
		t.Error("lookup variable fail.")
	}
}
