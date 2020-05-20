package main

import (
	"strings"
	"testing"
)

func TestSublimeContains(t *testing.T) {
	tests := []struct {
		text   string
		substr string
		pass   bool
	}{
		{"hello", "lo", true},
		{"abcdefg", "cf", true},
		{"abcdefg", "a", true},
		{"abcdefg", "b", true},
		{"abcdefg", "cfa", false},
		{"abcdefg", "aa", false},
		{"世界", "a", false},
		{"Hello 世界", "界", true},
		{"Hello 世界", "elo", true},
	}
	for _, v := range tests {
		res := SublimeContains(v.text, v.substr)
		if res != v.pass {
			t.Fatalf("Failed: %v - res:%v", v, res)
		}
	}
}

func TestGetFileExt(t *testing.T) {
	ext:=GetFileExt("ppt\\展厅.ppt")
	ext=strings.ToLower(ext)
	if ext!=".ppt"{
		t.Fatalf("Failed:%v - res:%v",ext,"ppt")
	}
}