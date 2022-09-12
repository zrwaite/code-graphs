package wakatime

import (
	"math"

	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/models"
)

func CreateSlice(s *svg.SVG, start float64, size float64, colour string) {
	end := ((start+size)/100)*2*math.Pi - math.Pi/2
	start = (start/100)*2*math.Pi - math.Pi/2
	var radius float64 = 400
	middleX := 480
	middleY := 460
	startX := middleX + int(math.Cos(start)*radius)
	startY := middleY + int(math.Sin(start)*radius)

	endX := middleX + int(math.Cos(end)*radius)
	endY := middleY + int(math.Sin(end)*radius)

	var large bool
	if size > 50 {
		large = true
	}
	s.Arc(startX, startY, int(radius), int(radius), 0, large, true, endX, endY, "fill:"+colour)
	s.Polygon([]int{middleX, startX, endX}, []int{middleY, startY, endY}, "fill:"+colour)
}

func CreatePiGraph(s *svg.SVG, languages models.Languages) {
	s.Circle(480, 460, 400, "fill:none;stroke:white;stroke-width:5")
	totalAngle := 0.0
	for _, language := range languages.Languages {
		if language.Percent > 0.2 {
			CreateSlice(s, totalAngle, language.Percent, language.Colour)
			totalAngle += language.Percent
		}
	}
}
