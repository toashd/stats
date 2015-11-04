package stats

import "errors"

// Pearson calculates the Pearson product-moment correlation coefficient between two variables.
func Pearson(data1, data2 Float64Data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return 0, errors.New("Input data must not be empty")
	}

	if l1 != l2 {
		return 0, errors.New("Input data must be same length")
	}

	sdev1, _ := StandardDeviationPopulation(data1)
	sdev2, _ := StandardDeviationPopulation(data2)

	if sdev1 == 0 || sdev2 == 0 {
		return 0, nil
	}

	covp, _ := CovariancePopulation(data1, data2)
	return covp / (sdev1 * sdev2), nil
}