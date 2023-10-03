package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// DebugJson Using for development purpose only. Printing as like as JSON object, more ez to debuging & reading object data.
func DebugJson(value interface{}) {
	fmt.Println(reflect.TypeOf(value).String())
	prettyJSON, _ := json.MarshalIndent(value, "", "    ")
	fmt.Printf("%s\n", string(prettyJSON))
}
