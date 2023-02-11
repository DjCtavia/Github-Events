package main

import (
	"fmt"
	"strings"
)

type dataHook struct {
	ID      string
	hookURL string
}

type hookURLFlag []string

func (hookFlag *hookURLFlag) String() string {
	return strings.Join(*hookFlag, ",")
}

func (hookFlag *hookURLFlag) Set(value string) error {
	*hookFlag = append(*hookFlag, value)
	return nil
}

func (hookFlag *hookURLFlag) GetArrayOfDataHook() ([]*dataHook, error) {
	var dataHookArray []*dataHook
	for index, hook := range *hookFlag {
		dataArray := strings.Split(hook, `@`)

		if len(dataArray) != 2 {
			return nil, fmt.Errorf("flag index %d couldn't separate ID and hookURL on: %s. use the format ID@HookURL", index, hook)
		}
		dataHookArray = append(dataHookArray, &dataHook{dataArray[0], dataArray[1]})
	}
	return dataHookArray, nil
}
