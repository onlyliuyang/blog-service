package test

import (
	"github.com/blog-service/pkg/util"
	"testing"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func TestStructToMap(t *testing.T) {
	//m := make(map[string]interface{})
	person := Person{
		Name:    "zhangsan",
		Address: "北京海淀",
	}

	//b, _ := json.Marshal(person)
	//_ = json.Unmarshal(b, &m)
	util.StructToMap(person)
}
