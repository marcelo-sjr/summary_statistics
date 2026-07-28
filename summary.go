package summary

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type integer interface {
	~int | ~int32 | ~int64 | ~int8 | ~int16
}

type float interface {
	~float32 | ~float64
}

type IntSummary struct {
	Count int
	Sum   int
	Avg   float64
	Min   int
	Max   int
}

type FloatSummary struct {
	Count int
	Sum   float64
	Avg   float64
	Min   float64
	Max   float64
}

// Ints computes summary statistics for a collection of integer values.
// the average is returned as a float64.
func Ints[T integer](t ...T) IntSummary {
	if len(t) == 0 {
		return IntSummary{}
	}

	var sum T
	min := t[0]
	max := t[0]

	for _, value := range t {
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
		Sum:   int(sum),
		Avg:   float64(sum) / float64(len(t)),
		Min:   int(min),
		Max:   int(max),
	}
}

// Float computes summary statistics for a collection of floating-point values.
// If the input is float32, the returned statistics are promoted to
// float64.
func Floats[T float](t ...T) FloatSummary {
	if len(t) == 0 {
		return FloatSummary{}
	}
	var sum T
	min, max := t[0], t[0]
	for _, value := range t {
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
		Sum:   float64(sum),
		Avg:   float64(sum) / float64(len(t)),
		Min:   float64(min),
		Max:   float64(max),
	}
}

// StreamInts reads integers from r and computes summary statistics.
//
// Each line is parsed as an integer using a bufio.Scanner.
// Lines that cannot be parsed are ignored.
//
// StreamInt does not close r; the caller is responsible for closing it,
// if necessary.
//
// If the scanner encounters an I/O error, an empty IntSummary and
// the corresponding error are returned.
func StreamInts(r io.Reader) (IntSummary, error) {
	scanner := bufio.NewScanner(r)

	var (
		count int
		sum   int
		min   int
		max   int
	)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
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

// StreamFloats reads floating-point numbers from r and computes summary
// statistics.
//
// Each line is parsed using strconv.ParseFloat with the provided bitSize.
// Lines that cannot be parsed are ignored.
//
// StreamFloat does not close r; the caller is responsible for closing it,
// if necessary.
//
// If the scanner encounters an I/O error, an empty FloatSummary and the
// corresponding error are returned.
func StreamFloats(r io.Reader, bitSize int) (FloatSummary, error) {
	if bitSize != 32 && bitSize != 64 {
		return FloatSummary{}, errors.New("invalid bitSize")
	}

	scanner := bufio.NewScanner(r)

	var (
		count int
		sum   float64
		min   float64
		max   float64
	)

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), bitSize)
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
