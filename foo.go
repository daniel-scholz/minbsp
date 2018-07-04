package main

//#include "sum.h"
//#include <windows.h>
//#include <winnt.h>
import "C"

func main() {
	//C.f(2, C.DEF)
	_ = C.PRAGMA_DEPRECATED_DDK
	f()
}
