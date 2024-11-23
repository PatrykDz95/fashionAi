package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"fasion.ai/server/internal/domain/recommendation"
)

func GetChatGPTResponse(prompt string) ([]recommendation.Item, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"

	reqBody := map[string]interface{}{
		"model": "gpt-4o-mini", // or gpt-3.5-turbo
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"max_tokens": 500,
	}

	jsonReq, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//log.Printf("Raw Response Body: %s", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	log.Printf("Raw message: %s", string(message))

	message, err = cleanUpMessage(message)

	var clothingItems []recommendation.Item
	if err := json.Unmarshal([]byte(message), &clothingItems); err != nil {
		return nil, errors.New("failed to parse JSON: " + err.Error())
	}

	return clothingItems, nil
}

func cleanUpMessage(rawResponse string) (string, error) {
	rawResponse = strings.TrimPrefix(rawResponse, "```json")
	rawResponse = strings.TrimSuffix(rawResponse, "```")

	return rawResponse, nil
}
