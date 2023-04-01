package sample_go_json2excel

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"strings"
	"testing"
)

func TestUserExcelData_ParseJSONCase1(t *testing.T) {
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
				j: strings.NewReader(`{"users":[{"name":"foo","age":20,"profile":"bar"},{"name":"foo2","age":20,"profile":"bar2"}]}`),
			},
			wantErr: false,
			wantData: &UserList{
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
			},
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
