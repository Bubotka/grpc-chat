package chat_gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChatGptClient struct {
	ApiKey      string
	OrgID       string
	ApiEndpoint string
}

func NewChatGptClient(apiKey string, orgID string, apiEndpoint string) *ChatGptClient {
	return &ChatGptClient{
		ApiKey:      apiKey,
		OrgID:       orgID,
		ApiEndpoint: apiEndpoint,
	}
}

func (c *ChatGptClient) PromptChatGPT(messages []RequestMessage) (string, error) {
	client := &http.Client{}

	// request payload
	payload := &ChatGptRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
		Stream:   false,
	}

	// convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal payload: %v", err)
		return "", err
	}
	log.Println("payload: ", string(payloadJSON))

	// create new HTTP request
	req, err := http.NewRequest("POST", c.ApiEndpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Printf("failed to create HTTP request: %v", err)
		return "", err
	}

	// set headers and authorization
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	req.Header.Set("OpenAI-Organization", c.OrgID)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to send HTTP request: %v", err)
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return "", err
	}

	// parse response
	var response ChatGptResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("failed to unmarshal JSON response: %v", err)
		var data map[string]interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			log.Printf("found error in JSON response: %v", err2)
		}
		return "", err
	}

	// extract the first choice from the array
	var ret string
	for _, choice := range response.Choices {
		log.Printf("finish reason is: %s", choice.FinishReason)
		ret = fmt.Sprintf("%s %s", ret, choice.Message.Content)
	}

	return ret, nil
}
