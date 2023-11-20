package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 20
	var b int
	var c = 4
	d := 5

	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d))

	var e = "Hello"
	var f string

	fmt.Println(e, f)
	fmt.Println(reflect.TypeOf(e), reflect.TypeOf(f))

	convertType()
	convertType2()
	typeCheck()
}

/*
여러가지 변수 선언법
-> var a int = 10
   var a int (초기 값을 적지 않을 경우 default 값이 선언된다. ex. int -> 0)
   var a = 10 (타입을 적지 않을 경우 default 타입이 선언된다. ex. 10 -> int)
   a := 10

여러가지 타입의 기본값
-> 모든 정수 타입: 0
   무든 실수 타입: 0.0
   불리언: false
   문자열: ""
   그 외: nil(정의되지 않은 메모리 주소를 나타내는 Go의 키워드)
*/

func convertType() {
	a := 3 // int
	var b float64 = 3.5
	var c int = int(b)
	// d := a * b
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e) // 64비트 컴퓨터에서 int는 int64와 같다고 했는데 왜 곱하기가 안될까? -> Go의 타입 체크가 철저하기 때문

	fmt.Println(a, b, c, d, e, f)
}

/*
타입 변환
-> Go 에서 연산의 각 항목의 타입은 "반드시" 같아야 한다.
*/

func convertType2() {
	var a int16 = 3456   // a는 int16타입 - 2바이트 정수
	var b int8 = int8(a) // int16 -> int8로 변환한 값

	// 큰 size의 값을 작은 size의 타입으로 변환하려고 해서 버려지는 값이 생김
	// 0000110110000000 -> int16
	//   버려짐   10000000 -> int8
	fmt.Println(a, b)
}

var g int = 10

func convertType3() {
	var m int = 20
	{
		var s int = 50
		fmt.Println(m, s, g)
	}
	// m = s + 20
}

/*
기본적으로 변수의 범위는 중괄호 안에서만 유효하다.
즉, g > m > s 순으로 쓰일수 있는 범위가 넓다.
g -> 패키지 전역 변수
m -> 함수에 속한 변수
s -> 중괄호 안에서만
*/

func typeCheck() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a, b, c, d)
}

/*
실질적으로 c와 d의 값은 실제 결과랑 다르다.
-> c는 0.3의 오차가 생긴 반면 d는 그 3배인 1의 오차가 생긴다
   즉, 실수연산에서는 오차를 항상 감안하고 줄일수 있는 방법을 강구해야함. 계산할수록 오차가 커진다.
*/
