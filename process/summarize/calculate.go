package main

import "sort"

// Summary is a 5-number-summary with the number of instances included.
type Summary struct {
	N       int
	Minimum float64
	Q1      float64
	Median  float64
	Q3      float64
	Maximum float64
}

// Summaries is a collection of summaries of different data-types.
type Summaries map[string]map[string]Summary

func calculateSummaries(data map[string]map[string][]float64) Summaries {
	ss := make(map[string]map[string]Summary)
	for name, datas := range data {
		ss[name] = make(map[string]Summary)
		for field, data := range datas {
			ss[name][field] = calculateSummary(data)
		}
	}
	return ss
}

func calculateMedian(data []float64) float64 {
	var median float64
	if len(data)%2 == 1 {
		median = data[(len(data)-1)/2]
	}
	if len(data)%2 == 0 {
		median = (data[(len(data)/2)-1] + data[len(data)/2]) / 2
	}
	return median
}

func lowerHalf(data []float64) []float64 {
	var half []float64
	if len(data)%2 == 1 {
		half = data[:(len(data)-1)/2]
	}
	if len(data)%2 == 0 {
		half = data[:len(data)/2]
	}
	return half
}

func upperHalf(data []float64) []float64 {
	var half []float64
	if len(data)%2 == 1 {
		half = data[((len(data)-1)/2)+1:]
	}
	if len(data)%2 == 0 {
		half = data[len(data)/2:]
	}
	return half
}

func calculateSummary(data []float64) Summary {
	if len(data) == 0 {
		return Summary{}
	}
	if len(data) == 1 {
		return Summary{
			N:       1,
			Minimum: data[0],
			Q1:      data[0],
			Median:  data[0],
			Q3:      data[0],
			Maximum: data[0],
		}
	}

	sort.Float64s(data)

	return Summary{
		N:       len(data),
		Minimum: data[0],
		Q1:      calculateMedian(lowerHalf(data)),
		Median:  calculateMedian(data),
		Q3:      calculateMedian(upperHalf(data)),
		Maximum: data[len(data)-1],
	}
}
