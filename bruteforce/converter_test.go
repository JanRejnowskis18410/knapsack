package bruteforce

import (
	"reflect"
	"testing"
)

func TestDecToBin(t *testing.T) {
	type args struct {
		x    int
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "0",
			args: args{x: 0, size: 1},
			want: []byte{48},
		},
		{
			name: "1 with size 1",
			args: args{x: 1, size: 1},
			want: []byte{49},
		},
		{
			name: "1 with size 2",
			args: args{x: 1, size: 2},
			want: []byte{0, 49},
		},
		{
			name: "15 with size 4",
			args: args{x: 15, size: 4},
			want: []byte{49, 49, 49, 49},
		},
		{
			name: "15 with size 5",
			args: args{x: 15, size: 5},
			want: []byte{0, 49, 49, 49, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecToBin(tt.args.x, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
