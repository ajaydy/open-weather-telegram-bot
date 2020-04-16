package helpers

func KelvinToCelcius(k float64) float64 {

	if k == 0 {
		return 0
	}

	c := k - 273.15

	return c

}
