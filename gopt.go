package main;

import (
	"fmt"; 
	"encoding/json";
); 

type any=interface{};

type OpenAIMessage struct {
	Role string;
	Content string;
}

func (message *OpenAIMessage) ToMap() map[string]string {	
	messageMap := make(map[string]string);

	messageMap["role"] = message.Role;
	messageMap["content"] = message.Content; 

	return messageMap;
}

type OpenAIRequest struct {
	Model string;
	Messages []OpenAIMessage;
}

func (request *OpenAIRequest) ToMap() map[string]any {
	requestMap := make(map[string]any);
	var messages []map[string]string; 

	requestMap["model"] = request.Model;

	for i := 0; i < len(request.Messages); i++ {
		messages = append(messages, request.Messages[i].ToMap());
	}

	requestMap["messages"] = messages;

	return requestMap;
}

func (request *OpenAIRequest) ToJson() any {
	jsonBytes, err := json.Marshal(request.ToMap());

	if (err != nil) {
		return nil;
	}	

	return string(jsonBytes)
} 

type GoptConfig struct {
	APIKey string;
}

type GoptInstance struct {
	Config GoptConfig;
}

func (instance *GoptInstance) Prompt(prompt string) bool {
	fmt.Println(prompt);

	return true;
}

func NewGoptInstance(config GoptConfig) *GoptInstance {
	instance := new(GoptInstance);

	instance.Config = config

	return instance;
}

func main() {
	request := OpenAIRequest{Model: "gpt-4", Messages: []OpenAIMessage{ OpenAIMessage{Role: "user", Content: "Hello World!"} }}

	fmt.Println(request.ToJson());
} 



