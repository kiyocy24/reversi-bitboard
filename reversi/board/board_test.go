package board

import (
	"fmt"
	"github.com/kiyocy24/reversi-bitboard/reversi/direction"
	"log"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kiyocy24/reversi-bitboard/reversi/player"
)

func TestBoard_Play(t *testing.T) {
	type fields struct {
		black  uint64
		white  uint64
		player player.Player
		turn   int
	}
	type args struct {
		p     player.Player
		row   int
		col   int
		input uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Board
		wantErr bool
	}{
		{
			name: "初手",
			fields: fields{
				black:  E4 | D5,
				white:  D4 | E5,
				player: player.Black,
				turn:   1,
			},
			args: args{
				p:     player.Black,
				input: C4,
			},
			want: &Board{
				black:  C4 | D4 | E4 | D5,
				white:  E5,
				player: player.White,
				turn:   2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				black:  tt.fields.black,
				white:  tt.fields.white,
				player: tt.fields.player,
				turn:   tt.fields.turn,
			}

			err := b.Reverse(tt.args.p, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reverse() error = %v, wantErr %v", err, tt.wantErr)
			}

			opts := cmp.Options{
				cmp.AllowUnexported(Board{}),
			}
			if diff := cmp.Diff(tt.want, b, opts...); diff != "" {
				wantBlackStr := fmt.Sprintf("- 	black bit:  %b", tt.want.black)
				gotBlackStr := fmt.Sprintf("+ 	black bit:  %b", b.black)
				wantWhiteStr := fmt.Sprintf("- 	white bit:  %b", tt.want.white)
				gotWhiteStr := fmt.Sprintf("+ 	white bit:  %b", b.white)
				diff = strings.Join([]string{diff, wantBlackStr, gotBlackStr, wantWhiteStr, gotWhiteStr}, "\n")
				t.Errorf("Reverse() is mismatch (-want +got)%s\n", diff)
			}
		})
	}
}

func TestBoard_reverse(t *testing.T) {
	type fields struct {
		black  uint64
		white  uint64
		player player.Player
		turn   int
	}
	type args struct {
		put uint64
	}
	type want struct {
		b *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "3枚",
			fields: fields{
				black:  C5,
				white:  D5,
				player: player.Black,
				turn:   1,
			},
			args: args{put: E5},
			want: want{
				b: &Board{
					black:  C5 | D5 | E5,
					white:  0,
					player: player.White,
					turn:   2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				black:  tt.fields.black,
				white:  tt.fields.white,
				player: tt.fields.player,
				turn:   tt.fields.turn,
			}
			b.reverse(tt.args.put)

			opts := cmp.Options{
				cmp.AllowUnexported(Board{}),
			}

			if diff := cmp.Diff(tt.want.b, b, opts...); diff != "" {
				t.Errorf("reverse() is mismatch (-want +got)%s\n", diff)
				log.Printf("%b", tt.want.b.black)
				log.Printf("%b", b.black)
			}
		})
	}
}

func Test_transfer(t *testing.T) {
	type args struct {
		put uint64
		d   direction.Direction
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "上",
			args: args{
				put: D4,
				d:   direction.Up,
			},
			want: D3,
		},
		{
			name: "右上",
			args: args{
				put: D4,
				d:   direction.UpperRight,
			},
			want: E3,
		},
		{
			name: "右",
			args: args{
				put: D4,
				d:   direction.Right,
			},
			want: E4,
		},
		{
			name: "右下",
			args: args{
				put: D4,
				d:   direction.LowerRight,
			},
			want: E5,
		},
		{
			name: "下",
			args: args{
				put: D4,
				d:   direction.Low,
			},
			want: D5,
		},
		{
			name: "左下",
			args: args{
				put: D4,
				d:   direction.LowerLeft,
			},
			want: C5,
		},
		{
			name: "左",
			args: args{
				put: D4,
				d:   direction.Left,
			},
			want: C4,
		},
		{
			name: "左上",
			args: args{
				put: D4,
				d:   direction.UpperLeft,
			},
			want: C3,
		},
		{
			name: "最上行",
			args: args{
				put: D1,
				d:   direction.Up,
			},
			want: 0,
		},
		{
			name: "最右列",
			args: args{
				put: H4,
				d:   direction.Right,
			},
			want: 0,
		},
		{
			name: "最下行",
			args: args{
				put: D8,
				d:   direction.Low,
			},
			want: 0,
		},
		{
			name: "最左列",
			args: args{
				put: A4,
				d:   direction.Left,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transfer(tt.args.put, tt.args.d)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("reverse() is mismatch (-want +got)%s\n", diff)
			}
		})
	}
}

func TestBoard_LegalBoard(t *testing.T) {
	type fields struct {
		black    uint64
		white    uint64
		player   player.Player
		opposite player.Player
		turn     int
	}
	type want struct {
		bi uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    want
	}{
		{
			name: "初手",
			fields: fields{
				black:    E4 | D5,
				white:    D4 | E5,
				player:   player.Black,
				opposite: player.White,
				turn:     1,
			},
			wantErr: false,
			want: want{
				bi: D3 | C4 | F5 | E6,
			},
		},
		{
			name: "2手目",
			fields: fields{
				black:    E4 | D5 | E5 | F5,
				white:    D4,
				player:   player.White,
				opposite: player.Black,
				turn:     2,
			},
			wantErr: false,
			want: want{
				bi: F4 | D6 | F6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				black:    tt.fields.black,
				white:    tt.fields.white,
				player:   tt.fields.player,
				opposite: tt.fields.opposite,
				turn:     tt.fields.turn,
			}
			got := b.LegalBoard()
			if diff := cmp.Diff(tt.want.bi, got); diff != "" {
				t.Errorf("LegalBoard() is mismatch (-want +got)%s\n", diff)
				log.Println("0123456789012345678901234567890123456789012345678901234567890123")
				log.Printf("%064b", tt.want.bi)
				log.Printf("%064b", got)
			}
		})
	}
}
