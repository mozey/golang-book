// Package names match the folders they fall in. There are ways around it,
// but it's a lot easier if you stay within this pattern.
package math

func Min(xs []float64) float64 {
	min := xs[0]
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

func Max(xs []float64) float64 {
	min := xs[0]
	for _, x := range xs {
		if min < x {
			min = x
		}
	}
	return min
}

// Function names starting with a capital letter are visible to other packages
// and programs. Names starting with lower case are private to the package
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	if (len(xs) == 0) {
		return 0
	}
	return total / float64(len(xs))
}