package main

import (
	"strconv"
	"testing"
)

func Test_Unpack(t *testing.T) {
	tests := []struct {
		get  string
		want string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}
	for name, test := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			res, err := Unpack(test.get)
			if res != test.want {
				t.Errorf("got %s, want %s", res, test.want)
				t.Error(err)
			}
		})
	}
}
