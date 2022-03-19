package cond_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/syntax/cond"
	"github.com/snowmerak/generics-for-go/v2/types/result"
)

func ExampleIfCond() {
	a := cond.If(func(t int) bool {
		return t == 0
	}, func(t int) string {
		return "zero"
	}).ElseIf(func(i int) bool {
		return i == 1
	}, func(i int) string {
		return "one"
	}).Else(func(i int) string {
		return "other"
	})
	fmt.Println(a.Run(0))
	fmt.Println(a.Run(1))
	fmt.Println(a.Run(2))
	// Output:
	// zero
	// one
	// other
}

func ExampleSwitchCond() {
	a := cond.Switch[int, string]().Case(0, func(i int) string {
		return "zero"
	}).Case(1, func(i int) string {
		return "one"
	}).Default(func(i int) string {
		return "other"
	})
	fmt.Println(a.Run(0))
	fmt.Println(a.Run(1))
	fmt.Println(a.Run(2))
	// Output:
	// zero
	// one
	// other
}

func ExampleSwitchCond_result() {
	result := result.OK(100)

	resultCond := cond.Switch[bool, int]().Case(true, func(b bool) int {
		return result.Unwrap()
	}).Default(func(b bool) int {
		return -99
	})

	fmt.Println(resultCond.Run(result.IsOK()))
	// Output:
	// 100
}
