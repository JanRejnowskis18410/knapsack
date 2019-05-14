package main

import (
	"reflect"
	"testing"

	"github.com/sidletsky/knapsack/repository"
)

func Test_decToBin(t *testing.T) {
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
			if got := decToBin(tt.args.x, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPerfectKnapsack(t *testing.T) {
	itemsCase1 := []repository.Item{
		{Id: 1, Value: 10, Weight: 20},
		{Id: 2, Value: 15, Weight: 5},
		{Id: 3, Value: 20, Weight: 25},
		{Id: 4, Value: 30, Weight: 30},
	}
	itemsCase2 := []repository.Item{
		{Id: 1, Value: 15, Weight: 10},
		{Id: 2, Value: 20, Weight: 35},
		{Id: 3, Value: 10, Weight: 5},
		{Id: 4, Value: 25, Weight: 20},
	}

	type args struct {
		items            []repository.Item
		knapsackCapacity int
	}
	tests := []struct {
		name                string
		args                args
		wantPerfectKnapsack Knapsack
	}{
		{
			name:                "case 1",
			args:                args{itemsCase1, 40},
			wantPerfectKnapsack: Knapsack{TotalWeight: 35, TotalValue: 45, CharacteristicVector: []byte{0, 49, 48, 49}},
		},
		{
			name:                "case 2",
			args:                args{itemsCase2, 40},
			wantPerfectKnapsack: Knapsack{TotalWeight: 35, TotalValue: 50, CharacteristicVector: []byte{49, 48, 49, 49}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPerfectKnapsack := getPerfectKnapsack(tt.args.items, tt.args.knapsackCapacity); !reflect.DeepEqual(gotPerfectKnapsack, tt.wantPerfectKnapsack) {
				t.Errorf("getPerfectKnapsack() = %v, want %v", gotPerfectKnapsack, tt.wantPerfectKnapsack)
			}
		})
	}
}
