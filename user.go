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

type UserList struct {
	Users []*User `json:"users"`
}

// ParseJSONCase1 は、ユーザー一覧をまとめて1回でパースする。
func (u *UserList) ParseJSONCase1(j io.Reader) error {
	dec := json.NewDecoder(j)
	if err := dec.Decode(u); err != nil {
		return err
	}

	return nil
}

func (u *UserList) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(u); err != nil {
		return err
	}

	return nil
}
