package cfgreader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type KeyDict interface {
	Key(string) string
}

type SimpleFileKeyDict struct {
	keys map[string]string
}

func (dict *SimpleFileKeyDict) init(filename string) (*SimpleFileKeyDict, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, errors.New("Dict File is Not Exist")
	}

	ditcFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	defer ditcFile.Close()

	byteValue, _ := ioutil.ReadAll(ditcFile)
	var results map[string][]string
	json.Unmarshal([]byte(byteValue), &results)

	dict.keys = make(map[string]string)
	for k, vs := range results {
		// fmt.Println(k)
		// fmt.Println(vs)
		dict.keys[k] = k
		for _, v := range vs {
			dict.keys[v] = k
		}
	}

	// fmt.Println(dict.keys)

	return dict, nil
}

func (dict *SimpleFileKeyDict) Key(value string) string {
	return dict.keys[value]
}

func NewSimpleFileKeyDict(filename string) (*SimpleFileKeyDict, error) {
	dict := SimpleFileKeyDict{}
	return dict.init(filename)
}
