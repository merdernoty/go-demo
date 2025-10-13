package main

// import (
// 	"fmt"
// 	"reflect"
// )

type Point struct {
	X int
	Y int
}

type Person struct {
	Name    string
	Friends []string
}

// func main() {
// 	p1 := Point{1, 2}
// 	p2 := Point{1, 2}
// 	p3 := Point{3, 4}

// 	fmt.Printf("p1 == p2: %t\n", p1 == p2) // true
// 	fmt.Printf("p1 == p3: %t\n", p1 == p3) // false

// 	person1 := Person{Name: "Alice", Friends: []string{"Bob"}}
// 	person2 := Person{Name: "Alice", Friends: []string{"Bob"}}
	

// 	fmt.Printf("person1 == person2: %t\n", reflect.DeepEqual(person1,person2))
// 	fmt.Printf("Имена равны: %t\n", person1.Name == person2.Name)
// }
