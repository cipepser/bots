package util

import (
	"bufio"
	"encoding/json"
	"os"
)

func GetToken(filepath string, token interface{}) error {
	// TODO: pathの処理

	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	r := bufio.NewReaderSize(f, 4096)

	return json.NewDecoder(r).Decode(token)
}
