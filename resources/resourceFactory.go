package resources

import "fmt"

func GetResource(resourceName string) (iResource, error) {
	if resourceName == "redis" {
		return newRedisClient(), nil
	}
	if resourceName == "maverick" {
		return newRedisClient(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
