package sample_go_json2excel

import (
	"encoding/json"
	"fmt"
	"io"
)

var ErrInvalidJSON = fmt.Errorf("invalid json")

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

// ParseJSONCase2 は、ユーザー一覧を1人ずつパースする。
func (u *UserList) ParseJSONCase2(j io.Reader) error {
	dec := json.NewDecoder(j)
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if elem, ok := t.(string); ok {
			switch elem {
			case "users":
				if err := u.parseUsers(dec); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (u *UserList) parseUsers(dec *json.Decoder) error {
	t, err := dec.Token()
	if err != nil {
		return err
	}
	// ユーザーの集合は配列で入っているはずなのでそれをチェック
	if elem, ok := t.(json.Delim); !ok || elem != '[' {
		return ErrInvalidJSON
	}

	for dec.More() {
		var user User
		if err := dec.Decode(&user); err != nil {
			return err
		}

		u.Users = append(u.Users, &user)
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
