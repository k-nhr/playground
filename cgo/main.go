package main

/*
#include <stdlib.h>
#include <stdio.h>

char* hello(char *str)
{
	char *ptr = NULL;
	ptr = (char *)calloc(32, sizeof(char));
	if (ptr == NULL) {
		return NULL;
	}
	sprintf(ptr, "hello %s", str);
	printf("C: %s\n", ptr);

	return ptr;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str := C.CString("cgo")
	defer C.free(unsafe.Pointer(str))

	ptr := C.hello(str)
	fmt.Println("Go: " + C.GoString(ptr))
	defer C.free(unsafe.Pointer(ptr))

}
