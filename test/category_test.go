package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"
	"github.com/blog-service/pkg/util"
	"testing"
)

func init() {
	//err := initialization.SetupSetting("/Users/admin/go/src/github.com/go-programing-tour-book/blog-service/configs/app")
	//if err != nil {
	//	//t.Fatalf("init.setupSetting err: %v", err)
	//	fmt.Printf("init.setupSetting err: %v\n", err)
	//	return
	//}
	//
	//err = initialization.SetupRedis()
	//if err != nil {
	//	//t.Fatalf("init.setupRedis err: %v", err)
	//	fmt.Printf("init.setupRedis err: %v\n", err)
	//	return
	//}
}

func TestCreateCategory(t *testing.T) {
	params := []map[string]interface{}{
		{
			"name":      "臭豆腐",
			"level":     2,
			"status":    1,
			"parent_id": 1000005,
		},
		{
			"name":      "麻辣香锅",
			"level":     2,
			"status":    1,
			"parent_id": 1000005,
		},
		{
			"name":      "麻辣鸡丝",
			"level":     2,
			"status":    1,
			"parent_id": 1000005,
		},
		{
			"name":      "老婆饼",
			"level":     2,
			"status":    1,
			"parent_id": 1000005,
		},
		{
			"name":      "重庆小面",
			"level":     2,
			"status":    1,
			"parent_id": 1000005,
		},
		{
			"name":      "故宫",
			"level":     2,
			"status":    1,
			"parent_id": 1000006,
		},
		{
			"name":      "香山",
			"level":     2,
			"status":    1,
			"parent_id": 1000006,
		},
		{
			"name":      "北戴河",
			"level":     2,
			"status":    1,
			"parent_id": 1000006,
		},
		{
			"name":      "天之眼",
			"level":     2,
			"status":    1,
			"parent_id": 1000006,
		},
		{
			"name":      "东方明珠塔",
			"level":     2,
			"status":    1,
			"parent_id": 1000006,
		},
		//{
		//	"name":      "旅游",
		//	"level":     1,
		//	"status":    1,
		//	"parent_id": 0,
		//},
		//{
		//	"name":      "游戏",
		//	"level":     1,
		//	"status":    1,
		//	"parent_id": 0,
		//},
		//{
		//	"name":      "学术",
		//	"level":     1,
		//	"status":    1,
		//	"parent_id": 0,
		//},
		//{
		//	"name":      "教育",
		//	"level":     1,
		//	"status":    1,
		//	"parent_id": 0,
		//},
	}

	header := map[string]string{"Content-Type": "application/json"}
	for i := 0; i < len(params); i++ {
		param := params[i]
		body, err := util.Post(context.Background(), "http://127.0.0.1:8080/api/v1/categories", param, header)
		if err != nil {
			t.Fatalf("创建标签失败: %v", err)
		}

		var response app.CommonResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("创建标签失败: %v", err)
		}

		if response.Code != errcode.Success.Code() {
			t.Fatalf("创建标签失败: %s", response.Msg)
		}

		fmt.Println(string(body))
	}
}

func TestListCategory(t *testing.T) {
	body, err := util.Get(context.Background(), "http://127.0.0.1:8080/api/v1/categories", nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var response app.CommonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if response.Code != errcode.Success.Code() {
		t.Fatalf(response.Msg)
	}

	fmt.Println(string(body))
}

func TestUpdateCategory(t *testing.T) {
	params := map[string]interface{}{
		"id":     1000005,
		"name":   "美食王者",
		"status": 1,
		"level":  1,
	}
	header := map[string]string{"Content-Type": "application/json"}
	body, err := util.Put(context.Background(), "http://127.0.0.1:8080/api/v1/categories/"+"1000005", params, header)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var response app.CommonResponse
	_ = json.Unmarshal(body, &response)
	if response.Code != errcode.Success.Code() {
		t.Fatalf(response.Msg)
	}
	fmt.Println(response.Msg)
}
