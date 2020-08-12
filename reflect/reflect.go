package main

import (
	"fmt"
	"reflect"
)

type testStructX struct {
	Str	string
	Num	int
}

func (t testStructX)GetStr() string{
	return t.Str
}

func (t testStructX)GetStr2() string{
	return "lol"
}

func (t testStructY)GetStr2() string{
	return "kek"
}

func (t testStructX)GetNum() int {
	return t.Num
}

type testStructY struct {
	Str	string
	Num	int
}

func (t testStructY)GetStr() string{
	return t.Str
}

func (t testStructY)GetNum() int {
	return t.Num
}

func IsValidMethod(x interface{}, y interface{}) (bool, error) {
	xx := reflect.ValueOf(x)
	if xx.Kind() != reflect.Struct {
		return false, fmt.Errorf("[ERROR] x not a struct")
	}
	yy := reflect.ValueOf(y)
	if yy.Kind() != reflect.Struct {
		return false, fmt.Errorf("[ERROR] y not a struct")
	}
	if xx.NumMethod() != yy.NumMethod() {
		return false, nil
	}
	MapX := make(map[string]string, xx.NumMethod())
	for i := 0; i < xx.NumMethod(); i++ {
		name := xx.Type().Method(i).Name
		MapX[name] = ""
	}
	for i := 0; i < xx.NumMethod(); i++ {
		_, ok := MapX[yy.Type().Method(i).Name]
		if ok != true {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	x := testStructX{"kek", 42}
	y := testStructY{"lel", 21}
	IsEqual, _ := IsValidMethod(x, y)
	fmt.Println(IsEqual)
}
