package py

import (
	"fmt"
	"reflect"
)

func Print(x ...interface{}) {
	fmt.Println(x, reflect.TypeOf(x[0]))
}

func Input() string {
	var x string
	fmt.Scanln(&x)
	return x
}

func Append(x *interface{},y) {
	x = append(x,y)
}
