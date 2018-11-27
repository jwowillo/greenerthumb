package main

type window struct {
	Timestamp int64
	A, B, C   float64
}

func slide(w window, v float64) window {
	w.A = w.B
	w.B = w.C
	w.C = v
	return w
}

func average(w window) float64 {
	return w.A*1/6 + w.B*2/3 + w.C*1/6
}
