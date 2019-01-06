package main

import (
	"math"

	"github.com/jwowillo/greenerthumb/process"
)

// Pair associates a message.Header and a value.
type Pair struct {
	Header process.Header
	Value  float64
}

func clean(
	data map[string]map[string][]Pair,
	limit int) map[string]map[string][]Pair {
	cleanedData := make(map[string]map[string][]Pair)
	for name, datas := range data {
		cleanedData[name] = make(map[string][]Pair)
		for field, data := range datas {
			cleanedData[name][field] = cleanMoreThanLimit(data, limit)
		}
	}
	return cleanedData
}

func cleanMoreThanLimit(data []Pair, limit int) []Pair {
	vs := values(data)
	avg := average(vs)
	sigma := calculateStdDeviation(vs, avg)
	var xs []Pair
	for _, x := range data {
		dist := math.Abs(x.Value-avg) / sigma
		// > to allow data that is the exact number of
		// standard deviations away.
		if dist > float64(limit) {
			continue
		}
		xs = append(xs, x)
	}
	return xs
}

func calculateStdDeviation(xs []float64, avg float64) float64 {
	for i := range xs {
		xs[i] = math.Pow(xs[i]-avg, 2)
	}
	return math.Sqrt(average(xs))
}

func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func values(ps []Pair) []float64 {
	xs := make([]float64, len(ps))
	for i := range ps {
		xs[i] = ps[i].Value
	}
	return xs
}
