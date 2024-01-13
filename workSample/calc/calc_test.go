package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a, b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "正の数の場合", args: args{a: 2, b: 3}, want: 5},
		{name: "負の数の場合", args: args{a: -2, b: -3}, want: -5},
		{name: "0の場合", args: args{a: 0, b: 0}, want: 0},
	}

	// テストケースを反復処理
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add 関数を呼び出し
			got := Add(tt.args.a, tt.args.b)

			// 結果が期待値と一致するか確認
			if got != tt.want {
				t.Errorf("Add() = %v, want = %v", got, tt.want)
			}
		})
	}
}
