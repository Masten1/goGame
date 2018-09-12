package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name string
	age int
	pos position
}

func whereIsBadGuy(guy badGuy)  {
	x := guy.pos.x
	y := guy.pos.y

	fmt.Println("(Position X:", x,",", "Position Y:", y,")")
}

func main()  {
	p := position{10.55, 20}
	fmt.Println(p)

	badGuyPosition := position{99.99, 88.88}
	b := badGuy{"jabba", 100, badGuyPosition}

	whereIsBadGuy(b)
}
