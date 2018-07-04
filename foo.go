package main

//#include "sum.h"
import "C"

func main() {
	C.mySum(2, C.DEF)
	mySum()
}
