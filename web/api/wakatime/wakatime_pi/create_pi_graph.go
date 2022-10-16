package wakatime_pi

import (
	"math"

	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/models"
)

func CreateSlice(s *svg.SVG, start float64, size float64, colour string, addUsername bool, username string) {
	var radius float64 = 400
	middleX := 480
	middleY := 460
	if addUsername {
		middleY = 480
		radius = 380
	}
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
	s.Line(startX, startY, endX, endY, "stroke:"+colour+";stroke-width:1")
}

func CreatePiGraph(s *svg.SVG, languages models.Languages, addUsername bool, username string) {
	totalAngle := 0.0
	sliceLanguages := append(languages.Languages, models.Language{
		Name:         "Other",
		Colour:       "white",
		TotalSeconds: languages.Other.TotalSeconds,
		Percent:      languages.Other.Percent,
	})
	for i, language := range sliceLanguages {
		if addUsername {
			s.Text(480, 70, username, "text-anchor:middle; font:35pt Menlo; fill:white")
		}
		if i == 0 && language.Percent < 45 {
			if addUsername {
				s.Circle(480, 480, 380, "fill:none;stroke:white;stroke-width:5")
			} else {
				s.Circle(480, 460, 400, "fill:none;stroke:white;stroke-width:5")
			}
		}
		if language.Percent > 0.2 {
			CreateSlice(s, totalAngle, language.Percent, language.Colour, addUsername, username)
			totalAngle += language.Percent
		}
	}
}
