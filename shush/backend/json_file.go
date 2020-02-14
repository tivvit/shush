package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFile struct {
	data map[string]string
}

func NewJsonFile(fileName string) *JsonFile {
	jf := JsonFile{}
	jf.data = map[string]string{}
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &jf.data)
	if err != nil {
		fmt.Println(err)
	}
	return &jf
}

func (jf JsonFile) Get(key string) (string, error) {
	if v, ok := jf.data[key]; ok {
		return v, nil
	} else {
		return "", errors.New("key not found")
	}
}

func (jf JsonFile) GetAll() (map[string]string, error) {
	return jf.data, nil
}
