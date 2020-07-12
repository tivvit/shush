package backend

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

type JsonFile struct {
	data *InMem
	fn string
}

func NewJsonFile(fileName string) (*JsonFile, error) {
	jf := JsonFile{
		data: NewInMem(),
		fn: fileName,
	}
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := map[string]data{}
	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		return nil, err
	}
	tn := time.Now()
	for k, v := range data {
		if v.Ttl == nil {
			err := jf.data.Set(k, v.Data, 0)
			if err != nil {
				log.Warn(err)
			}
		} else {
			ttl := (*v.Ttl).Sub(tn)
			err := jf.data.Set(k, v.Data, ttl)
			if err != nil {
				log.Warn(err)
			}
		}
	}
	return &jf, nil
}

func (jf JsonFile) Get(key string) (string, error) {
	return jf.data.Get(key)
}

func (jf JsonFile) GetAll() (map[string]string, error) {
	return jf.data.GetAll()
}

func (jf JsonFile) Set(key string, value string, ttl time.Duration) error {
	err := jf.data.Set(key, value, ttl)
	if err != nil {
		return err
	}
	return jf.write()
}

func (jf JsonFile) Close() error {
	return jf.data.Close()
}

func (jf JsonFile) write() error {
	v, err := json.Marshal(jf.data.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(jf.fn, v, 0655)
}
