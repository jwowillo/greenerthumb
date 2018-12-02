package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/jwowillo/greenerthumb/process"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to parse input.
	ReadInput = 1 << iota
)

func main() {
	var ec int

	errorHandler := func(err error) { fmt.Fprintln(os.Stderr, err) }

	fieldHandler := makeFieldHandler(
		epsilon,
		equal,
		lessThan, lessThanOrEqual,
		greaterThan, greaterThanOrEqual)
	err := process.Fields(os.Stdin, fieldHandler, errorHandler)
	if err != nil {
		errorHandler(err)
		ec |= ReadInput
	}

	os.Exit(ec)
}

func makeFieldHandler(epsilon, e, lt, lte, gt, gte float64) process.FieldHandler {
	return func(name string, ts int64, field string, value float64) {
		isGood := true

		if !math.IsNaN(e) {
			if compareWithEpsilon(value, e, epsilon) != 0 {
				isGood = false
			}
		}
		if !math.IsNaN(lt) {
			if compareWithEpsilon(value, lt, epsilon) >= 0 {
				isGood = false
			}
		}
		if !math.IsNaN(lte) {
			if compareWithEpsilon(value, lte, epsilon) > 0 {
				isGood = false
			}
		}
		if !math.IsNaN(gt) {
			if compareWithEpsilon(value, gt, epsilon) <= 0 {
				isGood = false
			}
		}
		if !math.IsNaN(gte) {
			if compareWithEpsilon(value, gte, epsilon) < 0 {
				isGood = false
			}
		}

		if !isGood {
			return
		}

		fmt.Println(process.Serialize(name, ts, field, value))
	}
}

// compareWithEpsilon returns -1 if a < b, 0 if a == b, and 1 if a > b within
// the bounds of the passed epsilon.
func compareWithEpsilon(a, b, epsilon float64) int {
	v := a - b
	if v < -epsilon {
		return -1
	}
	if v > epsilon {
		return 1
	}
	return 0
}

var (
	name  string
	field string
)

var (
	epsilon            float64
	equal              float64
	lessThan           float64
	lessThanOrEqual    float64
	greaterThan        float64
	greaterThanOrEqual float64
)

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("filter instances of data-types by specifying a list of")
		p("ANDing conditions in the set of less than or equal to,")
		p("less than, equal, greater than, and greater than or equal")
		p("to and filtering STDIN according to the conditions.")
		p("")
		p("An epsilon value for comparisons can also optionall be")
		p("passed. The system epsilon is used otherwise.")
		p("")
		p("An example is:")
		p("")
		p("    ./filter A 1 --lt 4 --gt 2")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 0, \"1\": 1}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 2}")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 2, \"1\": 3}")
		p("    {\"Name\": \"A\", \"Timestamp\": 2, \"1\": 3}")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 3, \"1\": 4}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 4, \"1\": 5}")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")

		os.Exit(2)
	}

	flag.Float64Var(
		&epsilon,
		"epsilon",
		math.SmallestNonzeroFloat64,
		"epsilon for comparisons")
	flag.Float64Var(
		&equal,
		"e",
		math.NaN(),
		"number values must be equal to")
	flag.Float64Var(
		&lessThan,
		"lt",
		math.NaN(),
		"number values must be less than")
	flag.Float64Var(
		&lessThanOrEqual,
		"lte",
		math.NaN(),
		"number values must be less than or equal to")
	flag.Float64Var(
		&greaterThan,
		"gt",
		math.NaN(),
		"number values must be greater than")
	flag.Float64Var(
		&greaterThanOrEqual,
		"gte",
		math.NaN(),
		"number values must be greater than or equal to")

	if len(os.Args) < 3 {
		flag.Usage()
	}

	flag.CommandLine.Parse(os.Args[3:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}

	name = os.Args[1]
	field = os.Args[2]
}
