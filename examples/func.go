package main

import (
	"log"
//	"string"
)

func sum( a int, b int ) int {

	return a + b
}


func swap( x string, y string ) (string, string) {
	return y,x
}

func main() {
	log.Printf("Sum:%d\n", sum(8,4))
	x := "ONE"
	y := "TWO"
	a,b := swap(x,y)
	log.Printf("%s, %s", a,b)
}
