package summary

import (
	"bufio"
	"io"
	"strconv"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type float interface {
	~float32 | ~float64
}

type IntSummary struct {
	Count int
	Sum   int64
	Avg   float64
	Min   int64
	Max   int64
}

type FloatSummary struct {
	Count int
	Sum   float64
	Avg   float64
	Min   float64
	Max   float64
}

// Ints computes summary statistics for a collection of integer values.
// The average is returned as a float64.
func Ints[T integer](t ...T) IntSummary {
	if len(t) == 0 {
		return IntSummary{}
	}

	var sum int64
	min := int64(t[0])
	max := int64(t[0])

	for _, v := range t {
		value := int64(v)

		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return IntSummary{
		Count: len(t),
		Sum:   sum,
		Avg:   float64(sum) / float64(len(t)),
		Min:   min,
		Max:   max,
	}
}

// Floats computes summary statistics for a collection of floating-point values.
// If the input is float32, the returned statistics are promoted to
// float64.
func Floats[T float](t ...T) FloatSummary {
	if len(t) == 0 {
		return FloatSummary{}
	}
	var sum float64
	min, max := float64(t[0]), float64(t[0])
	for _, v := range t {
		value := float64(v)
		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return FloatSummary{
		Count: len(t),
		Sum:   sum,
		Avg:   sum / float64(len(t)),
		Min:   min,
		Max:   max,
	}
}

// ReadInts reads integers from r and computes summary statistics.
//
// Input can be tokenized using the provided bufio.SplitFunc.
// If split is nil, bufio.ScanWords is used.
// Each token is parsed as an integer using strconv.Atoi.
// Tokens that cannot be parsed are ignored.
//
// ReadInts does not close r; the caller is responsible for closing it,
// if necessary.
//
// If the scanner encounters an I/O error, an empty IntSummary and the
// corresponding error are returned.
func ReadInts(r io.Reader, split bufio.SplitFunc) (IntSummary, error) {
	if split == nil {
		split = bufio.ScanWords
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(split)

	var (
		count int
		sum   int64
		min   int64
		max   int64
	)

	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			continue
		}

		if count == 0 {
			min = value
			max = value
		}

		count++
		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	if err := scanner.Err(); err != nil {
		return IntSummary{}, err
	}

	if count == 0 {
		return IntSummary{}, nil
	}

	return IntSummary{
		Count: count,
		Sum:   sum,
		Avg:   float64(sum) / float64(count),
		Min:   min,
		Max:   max,
	}, nil
}

// ReadFloats reads floating-point numbers from r and computes summary
// statistics.
//
// Input can be tokenized using the provided bufio.SplitFunc.
// If split is nil, bufio.ScanWords is used.
// Each token is parsed using strconv.ParseFloat with the bitSize=64.
// Tokens that cannot be parsed are ignored.
//
// ReadFloats does not close r; the caller is responsible for closing it,
// if necessary.
//
// If the scanner encounters an I/O error, an empty FloatSummary and the
// corresponding error are returned.
func ReadFloats(r io.Reader, split bufio.SplitFunc) (FloatSummary, error) {
	if split == nil {
		split = bufio.ScanWords
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(split)

	var (
		count int
		sum   float64
		min   float64
		max   float64
	)

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			continue
		}

		if count == 0 {
			min = value
			max = value
		}

		count++
		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	if err := scanner.Err(); err != nil {
		return FloatSummary{}, err
	}

	if count == 0 {
		return FloatSummary{}, nil
	}

	return FloatSummary{
		Count: count,
		Sum:   sum,
		Avg:   sum / float64(count),
		Min:   min,
		Max:   max,
	}, nil
}
