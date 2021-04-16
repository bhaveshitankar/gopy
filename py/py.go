package py

import (
	"fmt"
	"reflect"
)

func Print(x ...interface{}) {
	//v, ok := in.(map[string]*x)
	fmt.Println(x, reflect.TypeOf(x[0]))
}

func Input() string {
	var x string
	fmt.Scanln(&x)
	//fmt.Scanf("%a %b %c",&a,&b,&c)
	return x
}
