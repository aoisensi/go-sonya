package sonya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (s *Discord) get(path string, v any) error {
	req, _ := http.NewRequest("GET", s.url(path), nil)
	req.Header.Set("Authorization", s.authorization())
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sonya: %v", err)
	}
	defer resp.Body.Close()
	if 200 > resp.StatusCode || resp.StatusCode >= 300 {
		return s.errors(resp)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func (s *Discord) post(path string, vo any, vi any) error {
	data, _ := json.Marshal(vi)
	req, _ := http.NewRequest("POST", s.url(path), bytes.NewBuffer(data))
	req.Header.Set("Authorization", s.authorization())
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sonya: %v", err)
	}
	defer resp.Body.Close()
	if 200 > resp.StatusCode || resp.StatusCode >= 300 {
		return s.errors(resp)
	}
	return json.NewDecoder(resp.Body).Decode(vo)
}

func (s *Discord) patch(path string, vo any, vi any) error {
	data, _ := json.Marshal(vi)
	req, _ := http.NewRequest("PATCH", s.url(path), bytes.NewBuffer(data))
	req.Header.Set("Authorization", s.authorization())
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sonya: %v", err)
	}
	defer resp.Body.Close()
	if 200 > resp.StatusCode || resp.StatusCode >= 300 {
		return s.errors(resp)
	}
	return json.NewDecoder(resp.Body).Decode(vo)
}

func (s *Discord) delete(path string, vo any) error {
	fmt.Println(path)
	req, _ := http.NewRequest("DELETE", s.url(path), nil)
	req.Header.Set("Authorization", s.authorization())
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sonya: %v", err)
	}
	defer resp.Body.Close()
	if 200 > resp.StatusCode || resp.StatusCode >= 300 {
		return s.errors(resp)
	}
	if resp.StatusCode == 204 {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(vo)
}

func (s *Discord) errors(resp *http.Response) error {
	var err ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &err)
	err.Raw = string(body)
	return err
}

func (s *Discord) url(path string) string {
	return s.BaseURL + "/v" + strconv.Itoa(s.APIVersion) + "/" + path
}

func (s *Discord) authorization() string {
	if s.isBot {
		return "Bot " + s.token
	} else {
		return "Bearer " + s.token
	}
}
