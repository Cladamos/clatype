package components

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/charmbracelet/lipgloss"
)

func RenderPerformanceChart(wpm, accuracy int, wpmHistory, accuracyHistory []int, width int, duration time.Duration) string {
	if len(wpmHistory) == 0 || len(accuracyHistory) == 0 {
		return "No data to display"
	}

	chartWidth := width - 10
	if chartWidth < 40 {
		chartWidth = 40
	}
	if chartWidth > 80 {
		chartWidth = 80
	}
	chartHeight := 12

	maxWPM := float64(maxInt(wpmHistory))
	maxAccuracy := 100.0
	maxY := maxWPM
	if maxAccuracy > maxY {
		maxY = maxAccuracy
	}
	if maxY < 10 {
		maxY = 10
	}

	totalSeconds := float64(len(wpmHistory))

	chart := linechart.New(
		chartWidth,
		chartHeight,
		0,            // minX
		totalSeconds, // maxX
		0,            // minY
		maxY,         // maxY
		linechart.WithXLabelFormatter(func(i int, v float64) string {
			seconds := int(v)
			if seconds%5 == 0 {
				return fmt.Sprintf("%ds", seconds)
			}
			if (seconds + 1) == int(duration.Seconds()) {
				return fmt.Sprintf("%ds", seconds+1)
			}
			return ""
		}),
		linechart.WithAutoYRange(),
	)

	wpmStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	accuracyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))

	for i := 0; i < len(wpmHistory)-1; i++ {
		p1 := canvas.Float64Point{X: float64(i), Y: float64(wpmHistory[i])}
		p2 := canvas.Float64Point{X: float64(i + 1), Y: float64(wpmHistory[i+1])}
		chart.DrawBrailleLineWithStyle(p1, p2, wpmStyle)
	}

	for i := 0; i < len(accuracyHistory)-1; i++ {
		p1 := canvas.Float64Point{X: float64(i), Y: float64(accuracyHistory[i])}
		p2 := canvas.Float64Point{X: float64(i + 1), Y: float64(accuracyHistory[i+1])}
		chart.DrawBrailleLineWithStyle(p1, p2, accuracyStyle)
	}

	chart.DrawXYAxisAndLabel()

	legend := lipgloss.JoinHorizontal(lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(fmt.Sprintf("■ %d WPM (Max: %d)", wpm, maxInt(wpmHistory))),
		"  ",
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render(fmt.Sprintf("■ %d%% Accuracy", accuracy)),
	)

	return lipgloss.JoinVertical(lipgloss.Left,
		legend,
		chart.View(),
	)
}

func maxInt(values []int) int {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
