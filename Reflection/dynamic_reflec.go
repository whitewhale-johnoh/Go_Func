package main

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]

	if a.Kind() != b.Kind() { // compare the types of both a and b
		fmt.Println("Different Types.")
		return nil
	}

	switch a.Kind() { // return a type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Unint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}

}

func makeSum(ftpr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()    // receive argument which is vacant interface and get real value information
	v := reflect.MakeFunc(fn.Type(), sum) // put sum func and create func info

	fn.Set(v) //create and connect the func via setting func info 'v' into fn which is info of fptr value
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))
	fmt.Println(flaotSum(2.1, 3.5))
	fmt.Println(stringSum("Hello, ", "world!"))

}
