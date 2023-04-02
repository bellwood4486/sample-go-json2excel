package sample_go_json2excel

import "testing"

func TestExcelize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Excelize()
		})
	}
}

func TestExcelizeStream(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExcelizeStream()
		})
	}
}

func TestExcelizeUserList(t *testing.T) {
	type args struct {
		list *UserList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				list: userList,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExcelizeUserList(tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("ExcelizeUserList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
