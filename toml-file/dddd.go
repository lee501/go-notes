package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func main() {
	content, err := ioutil.ReadFile("./config.toml")
	fmt.Println(err)
	v := viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(content))
	var res map[string]interface{}
	v.Unmarshal(&res)
	fmt.Println(res["go-validator"].([]interface{})[0].(map[string]interface{})["source_name"])
	fmt.Println(res["version"])
}
