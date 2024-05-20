package hah

import (
	"math"
	"sort"
	"strconv"
)

func ModifyText(s []string) string {
	var res string
	Ave := calculateAverage(s)
	Med := calculateMedian(s)
	Var := calculateVariance(s)
	Stand := calculateStandard(s)
	res = "the Average is: " + Ave + "\nthe Median is: " + Med + "\nthe Variance is: " + Var + "\nthe Standard Deviation: " + Stand
	return res
}

func calculateAverage(s []string) string {
	var sum float64
	a := StoF(s)
	for i := 0; i <= len(a)-1; i++ {
		sum += a[i]
	}

	average := sum / float64(len(a))
	str := strconv.Itoa(int(average))
	return str
}

func calculateMedian(s []string) string {
	a := StoF(s)

	sort.Float64s(a)
	length := len(a)
	if length == 0 {
		return "0"
	}

	middle := length / 2
	if length%2 == 1 {
		return strconv.Itoa(int(a[middle]))
	}

	result := (a[middle-1] + a[middle]) / 2
	final := strconv.Itoa(int(result))
	return final
}

func calculateVariance(s []string) string {
	//calculate the mean
	average := calculateAverage(s)
	mean, err := strconv.ParseFloat(average, 64)
	if err != nil {
		return "Couldnt convert number in standard diviation"
	}
	//turn string to floats and append to a float array
	a := StoF(s)

	var variance float64

	for _, num := range a {
		diff := num - mean
		variance += math.Pow(diff, 2)
	}
	n := len(a)
	if n == 0 {
		return "0"
	}
	last := variance / float64(n-1)
	res := strconv.Itoa(int(last))
	return res
}

func calculateStandard(s []string) string {
	varianceS := calculateVariance(s)
	variance, err := strconv.ParseFloat(varianceS, 64)
	if err != nil {
		return "Couldnt convert number in standard diviation"
	}

	standardDeviation := math.Sqrt(variance)

	res := strconv.Itoa(int(standardDeviation))
	return res
}

func StoF(s []string) []float64 {
	var a []float64
	for _, str := range s {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return a
		}
		a = append(a, num)
	}
	return a
}
