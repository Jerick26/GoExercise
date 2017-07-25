package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func EmployeeByID(id int) *Employee { return new(Employee) }

func testEmployeeStruct() {
	var dilbert Employee
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Seninor " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}

type Point struct{ X, Y int }

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func testMapKey() {
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

type Circle struct {
	// Center Point
	Point
	Radius int
}

type Wheel struct {
	// Circle Circle
	Circle
	Spokes int
}

func testEmbeddingAndAnomymous() {
	var w Wheel
	// w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	w.X = 42
	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:20}
	// w.Circle.Point.X = 8
	// // w.Circle.Center.Y = 8
	// // w.Circle.Radius = 5
	// w.X = 8
	// w.Y = 8
	// w.Radius = 5
	// w.Spokes = 20
}

func main() {
	testEmployeeStruct()
	p := Point{1, 2}
	// pp := &Point{1, 2}
	// exactly equivalent to
	pp := new(Point)
	*pp = Point{1, 2}
	// test comparing
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"
	pp1 := &Point{3, 4}
	pp2 := &Point{3, 4}
	fmt.Println(pp1 == pp2) // "false"
	testMapKey()
	testEmbeddingAndAnomymous()
}
