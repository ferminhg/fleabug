package fleabug

import (
	"fmt"
	"reflect"
)

// formatter get object and formatt it
func formatter(object interface{}) string {
	xValue := reflect.ValueOf(object)
	xType := xValue.Type()
	logger := fmt.Sprintf("%v\t(%v)", xValue, xType)
	return logger
}
