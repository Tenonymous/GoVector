package main

import (
	"fmt"
	"vector/src/packages/vector"
)

type Person struct {
	Age int
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("Age: %d, Name: %s", p.Age, p.Name)
}


func main() {
	var my_vec vector.Vec// = make(vector.Vec, 2)
	my_vec2 := vector.Vec{1,2,3,4,5}

	fmt.Println(my_vec2.Size())
	fmt.Println(my_vec.Empty())
	fmt.Println(my_vec2)
	my_vec2.PushBack(10)
	fmt.Println(my_vec2)
	my_vec2.PopBack()
	fmt.Println(my_vec2)
	fmt.Println(my_vec2.Find(2))
	fmt.Println(my_vec2.Find(8))
	my_vec3 := vector.Vec{5,1,4,0,13,42,3,4,8,0}
	fmt.Println(my_vec3)
	my_vec3.Sort(func(i, j int) bool { return my_vec3[i].(int) < my_vec3[j].(int) })
	fmt.Println(my_vec3)
	vec_per := vector.Vec{Person{2, "tobi"}, Person{1, "Antek"}}
	fmt.Println(vec_per)
	fmt.Println(vec_per)
	fmt.Println(vec_per.Find(Person{2, "tobi"}))
}