package main

import "reflect"

// NewModelSlice create a slice of any type
// for example:
//
//	NewModelSlice(int(1)).Interface() -> []int{}
//	NewModelSlice(struct{}).Interface() -> []struct{}
func NewModelSlice(model any) reflect.Value {
	sliceType := reflect.SliceOf(
		reflect.New(reflect.Indirect(reflect.ValueOf(model)).Type()).Type(),
	)

	return reflect.MakeSlice(sliceType, 0, 0)
}

// NewEmptyModel create a new empty model
// for example:
//
//	NewEmptyModel(int(1)).Interface() -> int{}
//	NewEmptyModel(struct{}).Interface() -> struct{}
func NewEmptyModel(model any) reflect.Value {
	return reflect.New(reflect.Indirect(reflect.ValueOf(model)).Type())
}

// AnyToAnySlice convert any to []any
// if base on Slice or Array return slice of any
// for example:
//
//	AnyToAnySlice([]int{1, 2, 3}) -> []any{1, 2, 3}
//	AnyToAnySlice(1) -> []any{1}
func AnyToAnySlice(slice any) []any {
	if v, ok := slice.([]any); ok {
		return v
	}

	reval := reflect.Indirect(reflect.ValueOf(slice))
	if reval.Type().Kind() != reflect.Slice &&
		reval.Type().Kind() != reflect.Array {
		return []any{slice}
	}

	length := reval.Len()
	ret := make([]any, length)
	for i := 0; i < length; i++ {
		ret[i] = reval.Index(i).Interface()
	}
	return ret
}
