package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetChatGPTResponse(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"

	reqBody := map[string]interface{}{
		"model": "gpt-4", // or gpt-3.5-turbo
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"max_tokens": 150,
	}

	jsonReq, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// Safely extract the message content from the response
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices found in response")
	}

	messageMap, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected structure in choices")
	}

	message, ok := messageMap["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected structure in message")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("content is not a string")
	}

	return content, nil
}
