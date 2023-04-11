package user

import (
	"strings"
	"testing"

	"gopkg.in/macaron.v1"
)

func TestValidateLogin(t *testing.T) {
	type args struct {
		ctx *macaron.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试用户登录",
			args: args{
				ctx: &macaron.Context{},
			},
			want: "操作成功",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateLogin(tt.args.ctx); strings.Contains(got, tt.want) {
				t.Errorf("ValidateLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
