package main

import "fmt"

func addOne(num *int)  {
	*num = *num + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name string
	age int
	pos position
}

func whereIsBadGuy(guy *badGuy)  {
	x := guy.pos.x
	y := guy.pos.y

	fmt.Println("(Position X:", x,",", "Position Y:", y,")")
}

func main()  {
	x := 5
	var xPtr = &x
	addOne(xPtr)
	fmt.Println(x)

	p := position{55.5, 66.6}
	badGuy := badGuy{"Jabba", 33, p}
	whereIsBadGuy(&badGuy)
}