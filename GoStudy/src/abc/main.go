package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (s Student) SayHi() {
	fmt.Printf("Student: I'm Student, in %s school\n", s.school)
}

func (s Student) Sing(lyrics string) {
	fmt.Printf("Student: lala...%s\n", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//Employee重载Human的Sing方法
func (e Employee) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Human) Set(){
	e.name = "zhuxinquna"
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type StudentMen interface {
	SayHi()
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	//paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	//sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	//Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men
	var women People

	//i能存储Student
	i = &mike
	women = People(i)
	women.WomenSay()
	//fmt.Println("This is Mike, a Student:")
	i.SayHi()
	//i.Set()
	//i.SayHi()

	//fmt.Println(mike.name)
}
