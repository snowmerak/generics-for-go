package pipe

import (
	"reflect"
)

const errorTypeName = "error"

func Link[T any](funs ...interface{}) func(in ...interface{}) T {
	result := *new(T)
	resultType := reflect.TypeOf(result)
	wrapped := func(in ...interface{}) T {
		resultValue := reflect.New(resultType).Elem()
		values := make([]reflect.Value, len(in))
		for i, v := range in {
			values[i] = reflect.ValueOf(v)
		}
		for _, fun := range funs {
			values = reflect.ValueOf(fun).Call(values)
			funcType := reflect.TypeOf(fun)
			if funcType.Out(funcType.NumOut()-1).Name() == errorTypeName {
				lastValue := values[len(values)-1]
				if !lastValue.IsNil() {
					for i := 0; i < resultType.NumField(); i++ {
						field := resultType.Field(i)
						if field.Type.Name() == errorTypeName {
							resultValue.Field(i).Set(lastValue)
							return resultValue.Interface().(T)
						}
					}
					return result
				}
				values = values[:len(values)-1]
			}
		}
		out := make([]interface{}, len(values))
		for i, v := range values {
			out[i] = v.Interface()
		}
		if len(out) == 1 && resultType.Kind() == reflect.TypeOf(out[0]).Kind() {
			return out[0].(T)
		}
		for i := 0; i < resultType.NumField(); i++ {
			field := resultType.Field(i)
			for j := 0; j < len(out); j++ {
				if field.Type.Name() == reflect.TypeOf(out[j]).Name() {
					resultValue.Field(i).Set(reflect.ValueOf(out[j]))
					out = append(out[:j], out[j+1:]...)
					break
				}
			}
		}
		return resultValue.Interface().(T)
	}
	return wrapped
}
