package util

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Post(ctx context.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodPost, path, params, headers)
}

func Delete(ctx context.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodDelete, path, params, headers)
}

func Get(ctx context.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodGet, path, params, headers)
}

func Put(ctx context.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodPut, path, params, headers)
}

func doRequest(ctx context.Context, method string, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	bytes, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(bytes))
	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req = req.WithContext(ctx)
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buffer := make([]byte, 20480)
	length, _ := resp.Body.Read(buffer)
	return buffer[:length], nil
}
