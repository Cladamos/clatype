package components

import "time"

func CalculateWpm(target, input string, elapsed time.Duration) (wpm float64, accuracy float64) {
	if elapsed <= 0 {
		return 0, 0
	}

	tRunes := []rune(target)
	iRunes := []rune(input)

	var correctChars int
	totalTypedChars := len(iRunes)

	if totalTypedChars == 0 {
		return 0, 0
	}

	for i := 0; i < len(iRunes); i++ {
		if iRunes[i] == tRunes[i] {
			correctChars++
		}
	}

	accuracy = (float64(correctChars) / float64(totalTypedChars)) * 100.0

	words := float64(correctChars) / 5.0
	minutes := elapsed.Minutes()

	wpm = words / minutes

	return wpm, accuracy
}
