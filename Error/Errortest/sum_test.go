package calc

import "testing"

/*
func TestSum(t *testing.T){
    r := Sum(1,2)
    if r != 3{
        t.Error("결과값이 3이아닙니다. r=",r) //에러발생
    }
}*/

type TestData struct {
    argument1 int
    argument2 int
    result  int
}

var testdata = []TestData {
    {2,6,8},
    {-8,3,-5},
    {6,-6,0},
    {0,0,0},
}

func TestSum(t *testing.T) {
    for _, d := range testdata {
        r := Sum(d.argument1, d.argument2)
        if r != d.result {
            t.Errorf(
                "%d +%d의 결과값이 %d(이)가 아닙니다. r=%d",
                d.argument1,
                d.argument2,
                d.result,
                r,
            )
        }
    }
}