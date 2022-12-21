package chatGpt

import (
	"encoding/json"
	"os"

	"github.com/jeffcail/gorequest"
)

const BASEURL = "https://api.openai.com/v1/"

// Response
type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Completions Creates a completion for the provided prompt and parameters
// curl https://api.openai.com/v1/completions \
//  -H 'Content-Type: application/json' \
//  -H 'Authorization: Bearer chatGPT api key' \
//  -d '{
//  "model": "text-davinci-003",
//  "prompt": "Say this is a test",
//  "max_tokens": 7,
//  "temperature": 0
// }'
func Completions(msg string) (string, error) {
	h := buildRequestHeader()
	p := buildRequestBody(msg)
	url := BASEURL + "completions"
	b, err := gorequest.Post(url, h, p)
	if err != nil {
		return "", err
	}
	res := &Response{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return "", err
	}
	var reply string
	if len(res.Choices) > 0 {
		for _, v := range res.Choices {
			reply = v.Text
			break
		}
	}
	return reply, nil
}

func buildRequestHeader() (h map[string]string) {
	h = make(map[string]string)
	h["Content-Type"] = "application/json"
	h["Authorization"] = "Bearer " + os.Getenv("OpenApiKey")
	return
}

func buildRequestBody(msg string) (p map[string]interface{}) {
	p = make(map[string]interface{})
	p["model"] = "text-davinci-003"
	p["prompt"] = msg
	p["max_tokens"] = 2048
	p["temperature"] = 0.7
	p["top_p"] = 1
	p["frequency_penalty"] = 0
	p["presence_penalty"] = 0
	return
}
