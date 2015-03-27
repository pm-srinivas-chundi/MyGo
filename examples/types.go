package main

import (
	"log"
)

func main() {
	const format = "%T -> [%v] \n"
	var (
		f bool
		b byte
		i int
		d float64
		s string
		i64 int64
		z complex128
	)

	log.Printf(format, f, f)
	log.Printf(format, b, b)
	log.Printf(format, i, i)
	log.Printf(format, d, d)
	log.Printf(format, s, s)
	log.Printf(format, i64, i64)
	log.Printf(format, z, z)
}
