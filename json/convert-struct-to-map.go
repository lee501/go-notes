package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int	`json:"age"`
}

func main() {
	u := UserInfo{Name: "lee", Age: 35}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	var u1 UserInfo
	json.Unmarshal(b, &u1)

	var m map[string]interface{}
	json.Unmarshal(b, &m)
	fmt.Printf("%#v\n", m)
	for k, v := range m{
		fmt.Printf("key: %#v value: %#v\n", k, v)
	}

}

func ToMap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accept struct or struct point; got %T", v)
	}

	t := v.Type()
	for i:=0; i< v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
