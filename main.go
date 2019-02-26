package main

import (
	"fmt"
	_ "github.com/segmentio/kafka-go/lz4"
)

func breakingRecords(scores []int32) []int32 {
	if scores == nil || len(scores) == 0 {
		return []int32{0, 0}
	}
	min, max := scores[0], scores[0]
	var minScores, maxScores int32 = 0, 0
	scoresLen := len(scores)
	for i := 1; i < scoresLen; i++ {
		it := scores[i]
		if it < min {
			min = it
			minScores++
			continue
		}
		if it > max {
			max = it
			maxScores++
		}
	}
	return []int32{maxScores, minScores}
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
}
