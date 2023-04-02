package sample_go_json2excel

import (
	"io"
	"strings"
	"testing"
)

const (
	// ユーザーリストのJSON
	userListJSON = `
{
	"users":[
		{"name":"foo","age":20,"profile":"bar"},
		{"name":"foo2","age":20,"profile":"bar2"}
	]
}`
)

var (
	// ユーザーリストのJSONをパースした結果
	userList = &UserList{
		Users: []*User{
			{
				Name:    "foo",
				Age:     20,
				Profile: "bar",
			},
			{
				Name:    "foo2",
				Age:     20,
				Profile: "bar2",
			},
		},
	}
)

func TestToExcelCase1(t *testing.T) {
	type args struct {
		j io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				j: strings.NewReader(userListJSON),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToExcelCase1(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("ToExcelCase1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToExcelCase2(t *testing.T) {
	type args struct {
		j io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				j: strings.NewReader(userListJSON),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToExcelCase2(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("ToExcelCase2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToExcelCase3(t *testing.T) {
	type args struct {
		j io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				j: strings.NewReader(userListJSON),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToExcelCase3(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("ToExcelCase3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
