package main

import (
	"flag"
	"fmt"
	"time"

	"knapsack/bruteforce"
	"knapsack/repository"
)

func main() {
	file := flag.String("file", "data/12", "a string")
	fmt.Println("File:", *file)
	repo, err := repository.New(*file)
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	perfectKnapsack := bruteforce.GetPerfectKnapsack(repo.Items, repo.Capacity)
	duration := time.Since(startTime)
	fmt.Println("Time took:", duration)
	fmt.Println("Knapsack's capacity:", repo.Capacity)
	fmt.Println(perfectKnapsack.String())
	fmt.Println("Items:")
	for i, v := range perfectKnapsack.CharacteristicVector {
		// 49 represents 1, 48 represents 0
		if v == 49 {
			fmt.Println("id:", repo.Items[i].Id, "weight:", repo.Items[i].Weight, "value:", repo.Items[i].Value)
		}
	}
}
