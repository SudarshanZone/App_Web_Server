package models

import (
	"log"
	"strconv"
	"gopkg.in/ini.v1"
)

const MaxToken = 50

type EnvironmentManager struct {
	ConfigMap   map[string]string
	ServiceName string
}

func (Em *EnvironmentManager) InitProcessSpace(serviceName1 string, FileName string, ProcessName string) int {
	Em.ServiceName = serviceName1
	log.Printf("[%s] Initializing process space", Em.ServiceName)

	cfg, err := ini.Load(FileName)
	if err != nil {
		log.Printf("[%s] Error loading INI file: %s, Error: %v", Em.ServiceName, FileName, err)
		return -1
	}

	section, err := cfg.GetSection(ProcessName)
	if err != nil {
		log.Printf("[%s] Section '%s' not specified in INI file: %s, Error: %v", Em.ServiceName, ProcessName, FileName, err)
		return -1
	}

	Em.ConfigMap = make(map[string]string)
	for _, key := range section.Keys() {
		Em.ConfigMap[key.Name()] = key.String()
		log.Printf("[%s] Loaded key: %s, value: %s", Em.ServiceName, key.Name(), key.String())
	}

	if len(Em.ConfigMap) > MaxToken {
		log.Printf("[%s] Exceeding max token limit: MaxToken: %d, Count of tokens: %d", Em.ServiceName, MaxToken, len(Em.ConfigMap))
		return -1
	}

	return 0
}

func (Em *EnvironmentManager) GetProcessSpaceValue(token string) string {
	value, found := Em.ConfigMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", Em.ServiceName, token)
		return ""
	}
	return value
}

func (Em *EnvironmentManager) GetProcessSpaceValueAsInt(token string) int {
	valueStr, found := Em.ConfigMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", Em.ServiceName, token)
		return -1
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[%s] Failed to convert token '%s' value '%s' to integer: %v", Em.ServiceName, token, valueStr, err)
		return -1
	}
	return value
}

