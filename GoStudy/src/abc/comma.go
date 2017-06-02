package main

import (
	"strconv"
	"container/heap"
)

type Person struct {
	name string
	age  int
}

func (p Person) f1(a int) int {
	return a
}

func (p Person) f2(a int) int {
	return a
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

var pp Person

type Inter interface {
	pp.Interface
	f1()
	f2()
	String()
}