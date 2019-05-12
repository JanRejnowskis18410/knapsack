package main

import (
	"github.com/sidletsky/knapsack/repository"
	"reflect"
	"testing"
)

func Test_getPerfectKnapsack(t *testing.T) {
	itemsCase1 := []repository.Item{
		{Id: 1, Value: 10, Weight: 20},
		{Id: 2, Value: 15, Weight: 5},
		{Id: 3, Value: 20, Weight: 25},
		{Id: 4, Value: 30, Weight: 30},
	}
	knapsackcapacityCase1 := 40
	perfectKnapsackCase1 := getPerfectKnapsack(itemsCase1, knapsackcapacityCase1)
	if perfectKnapsackCase1.TotalWeight != 35 {
		t.Errorf("getPerfectKnapsack(...).TotalWeight = %d; want 35", perfectKnapsackCase1.TotalWeight)
	}
	if perfectKnapsackCase1.TotalValue != 45 {
		t.Errorf("getPerfectKnapsack(...).TotalValue = %d; want 45", perfectKnapsackCase1.TotalValue)
	}
	charVectorCase1 := []int{0, 1, 0, 1}
	if !reflect.DeepEqual(perfectKnapsackCase1.CharacteristicVector, charVectorCase1) {
		t.Errorf("getPerfectKnapsack(...).CharacteristicVector = %d; want {0, 1, 0, 1}", perfectKnapsackCase1.CharacteristicVector)
	}

	itemsCase2 := []repository.Item{
		{Id: 1, Value: 15, Weight: 10},
		{Id: 2, Value: 20, Weight: 35},
		{Id: 3, Value: 10, Weight: 5},
		{Id: 4, Value: 25, Weight: 20},
	}
	knapsackCapacityCase2 := 40
	perfectKnapsackCase2 := getPerfectKnapsack(itemsCase2, knapsackCapacityCase2)
	if perfectKnapsackCase2.TotalWeight != 35 {
		t.Errorf("getPerfectKnapsack(...).TotalWeight = %d; want 35", perfectKnapsackCase2.TotalWeight)
	}
	if perfectKnapsackCase2.TotalValue != 50 {
		t.Errorf("getPerfectKnapsack(...).TotalValue = %d; want 45", perfectKnapsackCase2.TotalValue)
	}
	charVectorCase2 := []int{1, 0, 1, 1}
	if !reflect.DeepEqual(perfectKnapsackCase2.CharacteristicVector, charVectorCase2) {
		t.Errorf("getPerfectKnapsack(...).CharacteristicVector = %d; want {1, 0, 1, 1}", perfectKnapsackCase2.CharacteristicVector)
	}
}
