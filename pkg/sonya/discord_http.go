package sonya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *Session) get(path string, v interface{}) error {
	req, _ := http.NewRequest("GET", s.url(path), nil)
	req.Header.Set("Authorization", s.authorization)
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func (s *Session) patch(path string, v interface{}, j interface{}) error {
	data, _ := json.Marshal(j)
	fmt.Println(string(data))
	req, _ := http.NewRequest("PATCH", s.url(path), bytes.NewBuffer(data))
	req.Header.Set("Authorization", s.authorization)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if 200 > resp.StatusCode || resp.StatusCode >= 300 {
		return s.errors(resp)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func (s *Session) errors(resp *http.Response) error {
	var err ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &err)
	err.Raw = string(body)
	return err
}

func (s *Session) url(path string) string {
	return s.BaseURL + "/" + s.APIVersion + "/" + path
}
