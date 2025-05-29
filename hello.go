package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func rectangle(x int, y int) (int, int) {

	area := x * y
	circumf := 2 * (x * y)
	return area, circumf
}

func printParity(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even. \n", x)
		return
	}
	fmt.Printf("%v is odd. \n", x)
}

func areaf(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area : [%v, %v]", x, y)
	}
	areaf := x * y
	return &areaf, nil
}

func  berea(m int,n int)(*int, error)  {
	if  m == 0 || n == 0 {
 return nil, fmt.Errorf("zero inputs: [%v,%v]",m,n)
	}
	berea := m * n
	return &berea, nil
}
func printmessage( id string, message string) {
	fmt.Printf("id: %v, message: %v\n", id, message)
}

type city struct{
	name string
	tempC float64
}


func newCity(n string) city{
	return city{name: n}
}

func (c *city) tempF()float64 {
	return (c.tempC * 9 / 5) +32
}



func main() {
c := newCity("London")
c.tempC = 15.5
tempf:= c.tempF()
	fmt.Printf("[%v]: tempC =%v; tempF=%v \n", c.name, c.tempC, tempf)
	fmt.Printf("City: %s, Temperature: %.2f°C\n", c.name, c.tempC)

	msg := "Hello, Go!"

	defer printmessage("1", msg)
	defer printmessage("2", msg)


	msg = "Hello, Go! - updated"
	fmt.Println(msg)


	fmt.Println("Hello, world!")

	x := 3
	y := 6

	// m := 2
	// n:=0
	// berea,err := berea(m,n)
	// if err!=nil{
	// 	panic(fmt.Sprintf("error: %v", err))
	// }

	// fmt.Printf("area(%v,%v) = %v; \n",m,n,*&berea)

	areaf, err := areaf(x, y)
	if err != nil {
		fmt.Print("error: ", err)
		return
	}

	fmt.Printf("area(%v,%v) = %v; \n", x, y, *areaf)

	sub := func(x int, y int) int {
		return x - y
	}
	printParity(4)
	printParity(5)
	// area ,circumf := rectangle(x,y)
	area, _ := rectangle(x, y)
	fmt.Printf("rectangle(%v,%v): area =%v;\n,", x, y, area)
	fmt.Printf(" add(%v,%v): %v\n", x, y, add(x, y))
	fmt.Printf(" sub(%v,%v): %v\n", x, y, sub(x, y))
}
