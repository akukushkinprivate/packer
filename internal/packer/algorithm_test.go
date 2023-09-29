package packer

import (
	"math"
	"reflect"
	"testing"
)

func Test_numberOfPacks(t *testing.T) {
	type args struct {
		items     uint64
		packSizes []uint64
	}
	tests := []struct {
		name string
		args args
		want map[uint64]uint64
	}{
		{
			name: "",
			args: args{
				items:     0,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				items:     1515,
				packSizes: []uint64{},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				items:     1,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{250: 1},
		},
		{
			name: "",
			args: args{
				items:     250,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{250: 1},
		},
		{
			name: "",
			args: args{
				items:     251,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{500: 1},
		},
		{
			name: "",
			args: args{
				items:     501,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{500: 1, 250: 1},
		},
		{
			name: "",
			args: args{
				items:     12001,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{5000: 2, 250: 1, 2000: 1},
		},
		{
			name: "",
			args: args{
				items:     999,
				packSizes: []uint64{250, 500, 1000, 2000, 5000},
			},
			want: map[uint64]uint64{1000: 1},
		},
		{
			name: "",
			args: args{
				items:     845,
				packSizes: []uint64{250, 500, 1000, 2000, 5000, 16, 13, 200},
			},
			want: map[uint64]uint64{13: 1, 16: 2, 200: 4},
		},
		{
			name: "",
			args: args{
				items:     math.MaxUint64,
				packSizes: []uint64{1, 2},
			},
			want: map[uint64]uint64{1: 1, 2: 9223372036854775807},
		},
		{
			name: "",
			args: args{
				items:     math.MaxUint64,
				packSizes: []uint64{1},
			},
			want: map[uint64]uint64{1: 18446744073709551615},
		},
		{
			name: "",
			args: args{
				items:     1,
				packSizes: []uint64{1},
			},
			want: map[uint64]uint64{1: 1},
		},
		{
			name: "",
			args: args{
				items:     252,
				packSizes: []uint64{250},
			},
			want: map[uint64]uint64{250: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPacks(tt.args.items, tt.args.packSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfPacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
