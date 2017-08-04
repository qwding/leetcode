package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	// fmt.Println("start:")
	start := -1
	insert := false
	for i := 0; i < len(intervals); i++ {
		fmt.Println("i:", i, start)
		if newInterval.Start > intervals[i].Start {
			if i == len(intervals)-1 {
				start = i + 1
				insert = true
				break
			}
			continue
		} else if newInterval.Start == intervals[i].Start {
			start = i
			break
		} else {
			if i == 0 {
				start = 0
				insert = true
			}
			if newInterval.Start > intervals[i-1].End {
				start = i
				insert = true
			} else {
				start = i - 1
			}
			break
		}
	}

	if insert {

		tmp := append([]Interval{}, intervals[start:]...)

		intervals = append(intervals[:start], newInterval)
		// fmt.Println("ins:", tmp, intervals)
		intervals = append(intervals, tmp...)

	} else {
		// fmt.Println("start:", start)
		if intervals[start].End >= newInterval.End {
			// fmt.Println("123:", intervals)
			return intervals
		} else {
			intervals[start].End = newInterval.End
		}
	}

	for i := start + 1; ; {
		if i >= len(intervals) {
			break
		}

		if intervals[start].End >= intervals[i].End {
			intervals = append(intervals[:i], intervals[i+1:]...)
			continue
		}
		if intervals[start].End >= intervals[i].Start {
			intervals[start].End = intervals[i].End
			intervals = append(intervals[:i], intervals[i+1:]...)
			continue
		}
		if intervals[start].End < intervals[i].Start {
			break
		}
	}

	// fmt.Println("int:", intervals)
	return intervals

}

func main() {
	arr := []Interval{}
	// arr = append(arr, Interval{1, 2})
	// arr = append(arr, Interval{3, 5})
	// arr = append(arr, Interval{6, 7})
	// arr = append(arr, Interval{8, 10})
	// arr = append(arr, Interval{12, 16})
	// in := Interval{4, 9}

	// arr = append(arr, Interval{1, 3})
	// arr = append(arr, Interval{6, 9})
	// in := Interval{4, 5}

	arr = append(arr, Interval{1, 5})
	in := Interval{2, 3}

	res := insert(arr, in)
	fmt.Println("res:", res)
}
