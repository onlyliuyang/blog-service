package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blog-service/internal/dao"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"
	"github.com/blog-service/pkg/util"
	"testing"
)

func TestCreateComment(t *testing.T) {
	params := []map[string]interface{}{
		{
			"article_id":   1000004,
			"commenter":    "阿科",
			"commenter_id": 10013,
			"content":      "你好，你叫什么名字",
			"parent_id":    0,
		},
		{
			"article_id":   1000004,
			"commenter":    "科大大",
			"commenter_id": 10014,
			"content":      "你好，我叫科大大",
			"parent_id":    0,
		},
		{
			"article_id":   1000004,
			"commenter":    "阿科",
			"commenter_id": 10013,
			"content":      "你会打篮球吗?",
			"parent_id":    0,
		},
		{
			"article_id":   1000004,
			"commenter":    "科大大",
			"commenter_id": 10014,
			"content":      "我不会，我只会唱唱歌",
			"parent_id":    0,
		},
	}

	var parentId interface{}
	for _, param := range params {
		header := map[string]string{"Content-Type": "application/json"}
		param["parent_id"] = parentId
		body, err := util.Post(context.Background(), "http://127.0.0.1:8080/api/v1/comments/", param, header)
		if err != nil {
			t.Fatalf(err.Error())
		}

		var response app.CommonResponse
		_ = json.Unmarshal(body, &response)
		if response.Code != errcode.Success.Code() {
			t.Fatalf(response.Msg, response.Details)
		}
		//fmt.Println(response.Data)
		d := response.Data.(map[string]interface{})
		parentId = d["id"]
	}
}

func TestGetComments(t *testing.T) {
	param := map[string]interface{}{"article_id": 1000004}
	header := map[string]string{"Content-Type": "application/json"}
	body, err := util.Get(context.Background(), "http://127.0.0.1:8080/api/v1/comments/", param, header)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var response app.CommonResponse
	_ = json.Unmarshal(body, &response)
	if response.Code != errcode.Success.Code() {
		t.Fatalf(response.Msg, response.Details)
	}
	//fmt.Println(response.Data)
	//d := response.Data.(map[string]interface{})
}

func TestUpdateComment(t *testing.T) {
	var d *dao.Dao
	err := d.UpdateComment(nil, nil)
	if err != nil {
		t.Fatalf("更新失败: %v", err)
	}
	fmt.Println("更新成功")
}
