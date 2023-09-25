package util;

import (
	"os"
);

func ReadJsonFile(path string) (any, error) {
	buffer, err := os.ReadFile(path);
	
	if (err != nil) {
		return nil, err;
	}

	return JsonToMap(string(buffer));
} 

func FileExists(path string) bool {
	_, err = os.Stat(path); 

	return err == nil;
}
