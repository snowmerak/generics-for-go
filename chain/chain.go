package chain

import (
	"errors"
	"reflect"

	"github.com/snowmerak/generics-for-go/result"
)

type Chain[T, R any] struct {
	list []any
}

func From[T, R any](list ...any) *result.Result[*Chain[T, R]] {
	if len(list) == 0 {
		return result.Failed[*Chain[T, R]](errors.New("list is empty"))
	}
	if len(list) == 1 {
		fun := reflect.TypeOf(list[0])
		if fun.Kind() != reflect.Func {
			return result.Failed[*Chain[T, R]](errors.New("list is not a function"))
		}
		if fun.In(0) != reflect.TypeOf(new(T)).Elem() {
			return result.Failed[*Chain[T, R]](errors.New("the function's parameter type is invalid"))
		}
		if fun.Out(0) != reflect.TypeOf(new(R)).Elem() {
			return result.Failed[*Chain[T, R]](errors.New("the function's return type is invalid"))
		}
		return result.Success(&Chain[T, R]{list})
	}
	in := reflect.TypeOf(list[0]).In(0)
	if in != reflect.TypeOf(new(T)).Elem() {
		return result.Failed[*Chain[T, R]](errors.New("function's parameter type is invalid"))
	}
	out := reflect.TypeOf(list[0]).Out(0)
	for i := 1; i < len(list); i++ {
		in = reflect.TypeOf(list[i]).In(0)
		if out != in {
			return result.Failed[*Chain[T, R]](errors.New("function's parameter type is invalid"))
		}
		out = reflect.TypeOf(list[i]).Out(0)
	}
	if out != reflect.TypeOf(new(R)).Elem() {
		return result.Failed[*Chain[T, R]](errors.New("function's return type is invalid"))
	}
	return result.Ok(&Chain[T, R]{list})
}

func (c *Chain[T, R]) Run(param T) R {
	paramValue := reflect.ValueOf(param)
	for _, fun := range c.list {
		funValue := reflect.ValueOf(fun)
		result := funValue.Call([]reflect.Value{paramValue})
		paramValue = result[0]
	}
	return paramValue.Interface().(R)
}
