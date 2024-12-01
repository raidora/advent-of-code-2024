package pretty

import (
	"fmt"
	"reflect"
	"strconv"
)

func PrintResult(result any, dayNum int, taskNum int) {
	var value string

	switch result.(type) {
	default:
		panic(fmt.Sprintf("Unable to pretty print value of type %v", reflect.TypeOf(result)))
	case int:
		value = strconv.Itoa(result.(int))
	case string:
		value = result.(string)
	}

	fmt.Printf("The result for DAY %d TASK %d:\r\n", dayNum, taskNum)
	fmt.Println(value)
}
