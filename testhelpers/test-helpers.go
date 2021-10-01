package testhelpers

import (
	"fmt"
	"reflect"
	"testing"
)

type TableDrivenTest struct {
	Name string
	Args []interface{}
	Want interface{}
}

func interfacesToValues(items []interface{}) []reflect.Value {
	var ret []reflect.Value

	for _, v := range items {
		ret = append(ret, reflect.ValueOf(v))
	}

	return ret
}

func valuesToInterfaces(items []reflect.Value) []interface{} {
	var ret []interface{}

	for _, v := range items {
		ret = append(ret, v.Interface())
	}

	return ret
}

func performTypeChecks(reflectedFunction reflect.Value, reflectedArgs []reflect.Value) error {
	if kind := reflectedFunction.Kind(); kind != reflect.Func {
		return fmt.Errorf("type error: Expected a function, got a %s instead", kind)
	}

	functionType := reflectedFunction.Type()

	if functionType.NumIn() != len(reflectedArgs) {
		return fmt.Errorf("type error: Function expects %d arguments, but received %d instead",
			functionType.NumIn(),
			len(reflectedArgs))
	}

	for i := 0; i < functionType.NumIn(); i++ {
		if !reflectedArgs[i].Type().AssignableTo(functionType.In(i)) {
			return fmt.Errorf("type error: At argument position %d: expected %s, but received %s",
				i,
				functionType.In(i),
				reflectedArgs[i].Type(),
			)
		}
	}

	return nil
}

func CallFunction(function interface{}, args []interface{}) ([]interface{}, error) {
	reflectedFunction := reflect.ValueOf(function)
	reflectedArgs := interfacesToValues(args)

	if err := performTypeChecks(reflectedFunction, reflectedArgs); err != nil {
		return nil, err
	}

	return valuesToInterfaces(reflectedFunction.Call(reflectedArgs)), nil
}

func RunTableDrivenTests(t *testing.T, functionUnderTest interface{}, tests []TableDrivenTest) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := CallFunction(functionUnderTest, tt.Args)

			if err != nil {
				t.Errorf("Error while calling function under test: %s", err)
				return
			}

			got := result[0]

			if got != tt.Want {
				t.Errorf("Got %v, want %v", got, tt.Want)
			}
		})
	}
}
