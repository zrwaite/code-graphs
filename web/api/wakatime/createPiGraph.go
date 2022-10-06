package wakatime

import (
	"math"

	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/models"
)

func CreateSlice(s *svg.SVG, start float64, size float64, colour string) {
	var radius float64 = 400
	middleX := 480
	middleY := 460
	if size > 50 {
		s.Arc(middleX, 60, int(radius), int(radius), 0, false, true, 480, 860, "fill:"+colour)
		size -= 50
		start += 50
	}
	end := ((start+size)/100)*2*math.Pi - math.Pi/2
	start = (start/100)*2*math.Pi - math.Pi/2

	startX := middleX + int(math.Cos(start)*radius)
	startY := middleY + int(math.Sin(start)*radius)

	endX := middleX + int(math.Cos(end)*radius)
	endY := middleY + int(math.Sin(end)*radius)
	s.Arc(startX, startY, int(radius), int(radius), 0, false, true, endX, endY, "fill:"+colour)
	s.Polygon([]int{middleX, startX, endX}, []int{middleY, startY, endY}, "fill:"+colour)
}

func CreatePiGraph(s *svg.SVG, languages models.Languages) {
	totalAngle := 0.0
	sliceLanguages := append(languages.Languages, models.Language{
		Name:         "Other",
		Colour:       "white",
		TotalSeconds: languages.Other.TotalSeconds,
		Percent:      languages.Other.Percent,
	})
	for i, language := range sliceLanguages {
		if i == 0 && language.Percent < 45 {
			s.Circle(480, 460, 400, "fill:none;stroke:white;stroke-width:5")
		}
		if language.Percent > 0.2 {
			CreateSlice(s, totalAngle, language.Percent, language.Colour)
			totalAngle += language.Percent
		}
	}
}
