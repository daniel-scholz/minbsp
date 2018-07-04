package main

//#include "sum.h"
//#include <windows.h>
//#include <winnt.h>
import "C"

func f() {
	//C.f(2, C.DEF)
	// call any constant from winnt.h to test for errors
	_ = C.PRAGMA_DEPRECATED_DDK

}
