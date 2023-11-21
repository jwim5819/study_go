package main

import (
	"fmt"
	"strings"
)

func main() {
	ex51()
}

func ex51() {
	var a int = 10
	var b int = 20
	var f float64 = 3279934123.115

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f)
	fmt.Println(strings.Repeat("=", 50))
}

/*
표준 출력 함수
-> Print(): 함수 입력값들을 출력
   Println(): 함수 입력값들을 출력하고 개행(입력값들 사이사이에 빈칸도 출력됨.)
   Printf(): 서식(format)에 맞도록 입력값들을 출력

Println()에서는 지수표현이고, Printf()에서는 소수점으로 표현되었다.
*/
