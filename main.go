package main;

import (
	"os"
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"gopt/util"
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
 
func (instance *GoptInstance) Prompt(prompt string) (any, error) {
	openAIRequest := OpenAIRequest{Model: "gpt-4", Messages: []OpenAIMessage{ OpenAIMessage{Role: "user", Content: prompt} }};
	requestStr, ok := openAIRequest.ToJson().(string);

	if (!ok) {
		return nil, errors.New("Failed to generate request payload.");
	}

	request, err := util.NewHttpRequest("POST", 
									fmt.Sprintf("%s/%s", instance.Config.BaseURL, "chat/completions"), 
									map[string]string{ "Content-Type": "application/json", "Authorization": fmt.Sprintf("Bearer %s", instance.Config.APIKey) },  
									requestStr);	
	
	if (err != nil) {
		fmt.Println(err);

		return false, err;
	}

	response, err := instance.HttpClient.Do(request.Request);

	if (err != nil) {
		return nil, err;
	}
	
	defer response.Body.Close();

	body, err := ioutil.ReadAll(response.Body); 

	if (err != nil)	{
		return nil, err;
	}
	
	bodyMapAny, err := util.JsonToMap(string(body));

	bodyMap := bodyMapAny.(map[string]any)

	errorMessage, ok := bodyMap["error"];
 
	if (ok) { // error exists
		return nil, errors.New(errorMessage.(map[string]any)["message"].(string));
	}

	return bodyMap["choices"].([]any)[0].(map[string]any)["message"].(map[string]any)["content"].(string), nil;
}


func main() {
	instance := GoptInstance{Config: GoptConfig{BaseURL: "https://api.openai.com/v1", APIKey: "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"}, HttpClient: &http.Client{}};

	prompt, err := util.CombineStrings(os.Args, " ", 1, len(os.Args) - 1);
	
	if (err != nil) {
		fmt.Println(err);

		return;
	}

	promptStr, ok := prompt.(string);
	
	if (!ok) {
		fmt.Println("Prompt typecast failed.");
	}

	s, err := instance.Prompt(promptStr); 
	
	if (err != nil) {
		fmt.Println(err);

		return;
	}
	
	result := s.(string);	

	fmt.Println(result);
} 


