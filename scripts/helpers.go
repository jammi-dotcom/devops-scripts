package devops_scripts

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

func GetLocalEnv() map[string]string {
	env := make(map[string]string)
	env["HOME"] = os.Getenv("HOME")
	env["GOROOT"] = os.Getenv("GOROOT")
	env["GOPATH"] = os.Getenv("GOPATH")
	return env
}

func GetDeploymentConfig() map[string]string {
	file, err := os.Open("config/deployment.yml")
	if err != nil {
		fmt.Println("Error opening deployment config file:", err)
		return map[string]string{}
	}
	defer file.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Println("Error reading deployment config file:", err)
		return map[string]string{}
	}
	config := strings.TrimSpace(buf.String())
	configMap := make(map[string]string)
	for _, line := range strings.Split(config, "\n") {
		if !strings.HasPrefix(line, "#") {
			keyValue := strings.SplitN(line, ":", 2)
			if len(keyValue) != 2 {
				continue
			}
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			configMap[key] = value
		}
	}
	return configMap
}