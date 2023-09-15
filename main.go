package main;

import (
	"fmt";
	"net/http";
	"encoding/json";
	"io/ioutil";
	"gopt/util";
); 

type any = interface{};

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

	return string(jsonBytes);
} 

type GoptConfig struct {
	BaseURL string;
	APIKey string;
}

type GoptInstance struct {
	Config GoptConfig;

	HttpClient *http.Client;
}
 
func (instance *GoptInstance) Prompt(prompt string) any {
	openAIRequest := OpenAIRequest{Model: "gpt-4", Messages: []OpenAIMessage{ OpenAIMessage{Role: "user", Content: prompt} }};
	requestStr, ok := openAIRequest.ToJson().(string);

	if (!ok) {
		return nil;
	}

	request, err := util.NewHttpRequest("POST", 
									fmt.Sprintf("%s/%s", instance.Config.BaseURL, "chat/completions"), 
									map[string]string{ "Content-Type": "application/json", "Authorization": fmt.Sprintf("Bearer %s", instance.Config.APIKey) },  
									requestStr);	
	
	if (err != nil) {
		fmt.Println(err);

		return false;
	}
	
	// request.Header.Add("Content-Type", "application/json");
	// request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", instance.Config.APIKey));

	response, err := instance.HttpClient.Do(request.Request);

	if (err != nil) {
		fmt.Println(err);

		return nil;
	}
	
	defer response.Body.Close();

	body, err := ioutil.ReadAll(response.Body); 

	if (err != nil)	{
		fmt.Println(err);
		return nil;
	}

	return string(body);
}

func main() {
	instance := GoptInstance{Config: GoptConfig{BaseURL: "https://api.openai.com/v1", APIKey: ""}, HttpClient: &http.Client{}};

	s, ok := instance.Prompt("Hello World!").(string);
	
	fmt.Println(s);
	fmt.Println(ok);
} 

