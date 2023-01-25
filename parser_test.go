package parser

import (
	"reflect"
	"testing"
)

func Test_initParser(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *parser
	}{
		{
			"test nil string",
			args{s: ""},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  [ string",
			args{s: "["},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  ] string",
			args{s: "["},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test  [] string",
			args{s: "[]"},
			&parser{
				s:    "",
				data: make([]string, 1, 1),
			}},

		{
			"test no space with comma string",
			args{s: "[1,2,3,4]"},
			&parser{
				s:    "1,2,3,4",
				data: []string{"1", "2", "3", "4"},
			}},

		{
			"test space with comma string",
			args{s: "[1 ,2,3 ,4]"},
			&parser{
				s:    "1,2,3,4",
				data: []string{"1", "2", "3", "4"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initParser(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
