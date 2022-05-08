package students

func Total(marks []float32) float32 {
	var total float32
	for _, mark := range marks {
		total += mark
	}
	return total
}

func Average(marks []float32) float32 {
	var total, average float32
	for _, mark := range marks {
		total += mark
	}
	subjectCount := len(marks)
	average = total / float32(subjectCount)
	return average
}
