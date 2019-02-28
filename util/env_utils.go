package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// LoadEnv method
func LoadEnv(config *string, name string) {
	v := os.Getenv(name)
	if v != "" {
		*config = v
	}
}

// LoadEnvInt method
func LoadEnvInt(config *int, key string) {
	v := os.Getenv(key)
	if v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			config = &i
		}
	}
}

// LoadEnvInt32 method
func LoadEnvInt32(config *int32, key string) {
	v := os.Getenv(key)
	if v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			i32 := int32(i)
			config = &i32
		}
	}
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
