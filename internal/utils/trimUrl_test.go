package utils

import "testing"

func TestTrimUrl(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name        string
		args        args
		wantTrimUrl string
	}{{
		"case1",
		args{
			"www.baidu.com",
		},
		"https://www.baidu.com",
	}, {
		"case2",
		args{
			"http://www.baidu.com",
		},
		"http://www.baidu.com",
	}, {
		"case3",
		args{
			"https://www.baidu.com",
		},
		"https://www.baidu.com",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTrimUrl := TrimUrl(tt.args.rawUrl); gotTrimUrl != tt.wantTrimUrl {
				t.Errorf("TrimUrl() = %v, want %v", gotTrimUrl, tt.wantTrimUrl)
			}
		})
	}
}
