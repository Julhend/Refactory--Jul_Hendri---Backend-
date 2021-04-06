package main

import (
	"fmt"
	"math"
	"sort"
)


type MeanMedian struct {
	numbers []float64
}

func main() {
	
	mmType := MeanMedian{
		numbers: []float64{9, 4, 2, 4, 1, 5, 3, 0},
	}

func (mm *MeanMedian) GetMinValue() float64 {
	sort.Float64s(mm.numbers)

	return mm.numbers[0]
}


func (mm *MeanMedian) GetMaxValue() float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	return mm.numbers[len(mm.numbers)-1]
}


func (mm *MeanMedian) CalcRangeValues() float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	return mm.GetMaxValue() - mm.GetMinValue()
}


func (mm *MeanMedian) CalcMean() float64 {
	total := 0.0

	for _, v := range mm.numbers {
		total += v
	}

	
	return math.Round(total / float64(len(mm.numbers)))
}


func (mm *MeanMedian) CalcMedian(n ...float64) float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	mNumber := len(mm.numbers) / 2

	if mm.IsOdd() {
		return mm.numbers[mNumber]
	}

	return (mm.numbers[mNumber-1] + mm.numbers[mNumber]) / 2
}

func (mm *MeanMedian) IsOdd() bool {
	if len(mm.numbers)%2 == 0 {
		return false
	}

	return true
}