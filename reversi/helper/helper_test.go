package helper

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kiyocy24/reversi-bitboard/reversi/board"
	"testing"
)

func TestCoordinateToBit(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "A1",
			args: args{row: 0, col: 0},
			want: board.A1,
		},
		{
			name: "H8",
			args: args{row: 7, col: 7},
			want: board.H8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoordinateToBit(tt.args.row, tt.args.col)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("CoordinateToBit() is mismatch (-want +got)%s\n", diff)
			}
		})
	}
}

func Test_BitToCoordinate(t *testing.T) {
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
			args: args{bi: board.A1},
			want: want{row: 0, col: 0},
		},
		{
			name: "7番地",
			args: args{bi: board.H1},
			want: want{row: 0, col: 7},
		},
		{
			name: "56番地",
			args: args{bi: board.A8},
			want: want{row: 7, col: 0},
		},
		{
			name: "63番地",
			args: args{bi: board.H8},
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
