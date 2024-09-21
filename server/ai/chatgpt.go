package ai

import (
	"bytes"
	"encoding/json"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return message, nil
}
