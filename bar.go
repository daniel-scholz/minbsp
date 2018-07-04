package main

//#include "sum.h"
import "C"

func mySum() {
	C.mySum(2, C.DEF)
}
