package day1

import "fmt"

/*

ReportRepair takes a slice of int, finds the two entries that sum 2020
and multiplies those two values together, and returns that value.
Returns Error if no values equal 2020

*/

func ReportRepair(report []int) (int, error) {
	if len(report) < 2 {
		return 0, fmt.Errorf("report must have at least two numbers")
	}
	for i, r := range report {
		for _, search := range report[i:] {
			if (r + search) == 2020 {
				return r * search, nil
			}
		}
	}
	return 0, fmt.Errorf("the report could not be fixed")
}
