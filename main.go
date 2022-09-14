package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {

	A := Record{"string value", -12.123, Secret{"ANurag ", "Gupta"}}
	r := reflect.ValueOf(A)

	fmt.Println("string Value:", r.String())

}
