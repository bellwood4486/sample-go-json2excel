package sample_go_json2excel

import (
	"github.com/google/go-cmp/cmp"
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

func TestUserList_ParseJSONCase1(t *testing.T) {
	type fields struct {
		Users []*User
	}
	type args struct {
		j io.Reader
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantData *UserList
	}{
		{
			name:   "success 2 users",
			fields: fields{},
			args: args{
				j: strings.NewReader(userListJSON),
			},
			wantErr:  false,
			wantData: userList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserList{
				Users: tt.fields.Users,
			}
			if err := u.ParseJSONCase1(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("ParseJSONCase1() error = %v, wantErr %v", err, tt.wantErr)
			}
			// go-cmpでdataを比較する
			if diff := cmp.Diff(u, tt.wantData); diff != "" {
				t.Errorf("ParseJSONCase1() diff = %v", diff)
			}
		})
	}
}

func TestUserList_ParseJSONCase2(t *testing.T) {
	type fields struct {
		Users []*User
	}
	type args struct {
		j io.Reader
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantData *UserList
	}{
		{
			name:   "success 2 users",
			fields: fields{},
			args: args{
				j: strings.NewReader(userListJSON),
			},
			wantErr:  false,
			wantData: userList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserList{
				Users: tt.fields.Users,
			}
			if err := u.ParseJSONCase2(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("ParseJSONCase2() error = %v, wantErr %v", err, tt.wantErr)
			}
			// go-cmpでdataを比較する
			if diff := cmp.Diff(u, tt.wantData); diff != "" {
				t.Errorf("ParseJSONCase2() diff = %v", diff)
			}
		})
	}
}
