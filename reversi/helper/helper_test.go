package helper

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_bitToRowAndCol(t *testing.T) {
	type args struct {
		bi uint64
	}
	type want struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0番地",
			args: args{bi: 1 << 0},
			want: want{row: 0, col: 0},
		},
		{
			name: "7番地",
			args: args{bi: 1 << 7},
			want: want{row: 0, col: 7},
		},
		{
			name: "56番地",
			args: args{bi: 1 << 56},
			want: want{row: 7, col: 0},
		},
		{
			name: "63番地",
			args: args{bi: 1 << 63},
			want: want{row: 7, col: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotCol := BitToCoordinate(tt.args.bi)
			if diff := cmp.Diff(tt.want.row, gotRow); diff != "" {
				t.Errorf("bitToRowAndCol() row is mismatch (-want +got)%s\n", diff)
			}
			if diff := cmp.Diff(tt.want.col, gotCol); diff != "" {
				t.Errorf("bitToRowAndCol() col is mismatch (-want +got)%s\n", diff)
			}
		})
	}
}
