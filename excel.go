package sample_go_json2excel

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
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

// ToExcelCase1 は以下の方式でExcelファイルを生成する。
//
// * 中間データの作成: あり
// * JSONのパース: bulk
// * Excelの書き込み: stream
func ToExcelCase1(j io.Reader) error {
	list, err := parseJSONBulk(j)
	if err != nil {
		return err
	}
	if err := createExcel(list); err != nil {
		return err
	}

	return nil
}

// ToExcelCase2 は以下の方式でExcelファイルを生成する。
//
// * 中間データの作成: あり
// * JSONのパース: stream
// * Excelの書き込み: stream
func ToExcelCase2(j io.Reader) error {
	list, err := parseJSONStream(j)
	if err != nil {
		return err
	}
	if err := createExcel(list); err != nil {
		return err
	}

	return nil
}

// ToExcelCase3 は以下の方式でExcelファイルを生成する。
//
// * 中間データの作成: なし
// * JSONのパース: stream
// * Excelの書き込み: stream
func ToExcelCase3(j io.Reader) error {
	return parseJSONAndCreateExcelStream(j)
}

func parseJSONBulk(j io.Reader) (*UserList, error) {
	dec := json.NewDecoder(j)
	var list UserList
	if err := dec.Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}

func parseJSONStream(j io.Reader) (*UserList, error) {
	dec := json.NewDecoder(j)
	var list UserList
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if elem, ok := t.(string); ok {
			switch elem {
			case "users":
				err := parseUsers(dec, func(u *User) error {
					list.Users = append(list.Users, u)
					return nil
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return &list, nil
}

func parseJSONAndCreateExcelStream(j io.Reader) error {
	f := excelize.NewFile()
	defer f.Close() // サンプルコードなのでエラーハンドリングは省略

	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}

	// ヘッダーを書き込む
	if err := sw.SetRow("A1",
		[]interface{}{
			excelize.Cell{Value: "Name"},
			excelize.Cell{Value: "Age"},
			excelize.Cell{Value: "Profile"},
		},
		excelize.RowOpts{Height: 45, Hidden: false}); err != nil {
		return err
	}

	// ユーザーを一人一行で書き込む
	rowID := 1
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
				if err := parseUsers(dec, func(u *User) error {
					rowID++
					cell, err := excelize.CoordinatesToCellName(1, rowID)
					if err != nil {
						return err
					}
					if err := sw.SetRow(cell, []interface{}{
						excelize.Cell{Value: u.Name},
						excelize.Cell{Value: u.Age},
						excelize.Cell{Value: u.Profile},
					}); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return err
				}
			}
		}
	}

	if err := sw.Flush(); err != nil {
		return err
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		return err
	}

	return nil
}

func parseUsers(dec *json.Decoder, fn func(u *User) error) error {
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
		if err := fn(&user); err != nil {
			return err
		}
	}

	return nil
}

func createExcel(list *UserList) error {
	f := excelize.NewFile()
	defer f.Close() // サンプルコードなのでエラーハンドリングは省略

	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}

	// ヘッダーを書き込む
	if err := sw.SetRow("A1",
		[]interface{}{
			excelize.Cell{Value: "Name"},
			excelize.Cell{Value: "Age"},
			excelize.Cell{Value: "Profile"},
		},
		excelize.RowOpts{Height: 45, Hidden: false}); err != nil {
		return err
	}

	// ユーザーを一人一行で書き込む
	rowID := 1
	for _, user := range list.Users {
		rowID++
		cell, err := excelize.CoordinatesToCellName(1, rowID)
		if err != nil {
			return err
		}
		if err := sw.SetRow(cell, []interface{}{
			excelize.Cell{Value: user.Name},
			excelize.Cell{Value: user.Age},
			excelize.Cell{Value: user.Profile},
		}); err != nil {
			return err
		}
	}

	if err := sw.Flush(); err != nil {
		return err
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		return err
	}

	return nil
}
