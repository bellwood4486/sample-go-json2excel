package sample_go_json2excel

import (
	"encoding/json"
	"io"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Profile string `json:"profile"`
}

type UserExcelData struct {
	Users []*User `json:"users"`
}

func (u *UserExcelData) ParseJSON(j io.Reader) error {
	dec := json.NewDecoder(j)
	if err := dec.Decode(u); err != nil {
		return err
	}

	return nil
}
