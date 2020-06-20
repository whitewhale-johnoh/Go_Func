//Atio - string which represents number changes to int type
//Itoa - int to string
//FormatBool - change bool into strings
//FormatFloat
//FormatInt
//FormatUint
//AppendBool - change bool into string and append it behind the slice
//AppendFloat
//AppendInt
//AppendUint
//ParseBool - change bool format string into bool type
//ParseFloat
//ParseInt
//ParseUintpackage

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var num1 int
	num1, err = strconv.Atoi("100")
	fmt.Println(num1, err)

	var num2 int
	num2, err = strconv.Atoi("10t")
	fmt.Println(num2, err)

	var s string
	s = strconv.Itoa(50)
	fmt.Println(s)

	var s1 string
	s1 = strconv.FormatBool(true)
	fmt.Println(s1)

	var s2 string
	s2 = strconv.FormatFloat(1.3, 'f', -1, 32)
	fmt.Println(s2)

	var s3 string
	s3 = strconv.FormatInt(-10, 10)
	fmt.Println(s3)

	var s4 string
	s4 = strconv.FormatUint(32, 16)
	fmt.Println(s4)

	var s5 []byte = make([]byte, 0)

	s5 = strconv.AppendBool(s5, true)
	fmt.Println(string(s5))

	s5 = strconv.AppendFloat(s5, 1.3, 'f', -1, 32)
	fmt.Println(string(s5))

	s5 = strconv.AppendInt(s5, -10, 10)
	fmt.Println(string(s5))

	s5 = strconv.AppendUint(s5, 32, 16)
	fmt.Println(string(s5))

	var err1 error

	var b1 bool
	b1, err1 = strconv.ParseBool("true")
	fmt.Println(b1, err1)

	var num3 float64
	num3, err1 = strconv.ParseFloat("1.3", 64)
	fmt.Println(num3, err1)

	var num4 int64
	num4, err1 = strconv.ParseInt("-10", 10, 32)
	fmt.Println(num4, err1)

	var num5 uint64
	num5, err1 = strconv.ParseUint("20", 16, 32)
	fmt.Println(num5, err1)

}
