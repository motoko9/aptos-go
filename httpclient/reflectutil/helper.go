package reflectutil

import "reflect"

func TypeOf(i interface{}) reflect.Type {
	return Indirect(ValueOf(i)).Type()
}

func ValueOf(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}

func Indirect(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

func KindOf(v interface{}) reflect.Kind {
	return TypeOf(v).Kind()
}
