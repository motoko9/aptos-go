package stringutil

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSplitWordsByCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "twoWords",
			args: args{str: "twoWords"},
			want: []string{"two", "Words"},
		},
		{
			name: "TwoWords",
			args: args{str: "TwoWords"},
			want: []string{"Two", "Words"},
		},
		{
			name: "oneword",
			args: args{str: "oneword"},
			want: []string{"oneword"},
		},
		{
			name: "ContinuousUPPER",
			args: args{str: "ContinuousUPPER"},
			want: []string{"Continuous", "UPPER"},
		},
		{
			name: "ContinuousUPPERLower",
			args: args{str: "ContinuousUPPERLower"},
			want: []string{"Continuous", "UPPER", "Lower"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitWordsByCamelCase(tt.args.str)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("SplitWordsByCamelCase() mismatch(-want,+got):\n%s", diff)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{str: "", want: true},
		{str: " ", want: true},
		{str: "\t\n", want: true},
		{str: "abc", want: false},
		{str: "  abc  ", want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("\"%s\" %v", tt.str, tt.want), func(t *testing.T) {
			if got := IsBlank(tt.str); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	type args struct {
		str string
		sub []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				str: "Abc",
				sub: []string{"ab", "d"},
			},
			want: true,
		},
		{
			args: args{
				str: "__hello_world",
				sub: []string{"LL"},
			},
			want: true,
		},
		{
			args: args{
				str: "LIST",
				sub: []string{"GG"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsIgnoreCase(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("ContainsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}