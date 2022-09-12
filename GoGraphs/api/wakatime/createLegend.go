package wakatime

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/models"
)

func CreateLegend(s *svg.SVG, languages models.Languages) {
	centerX := 1200

	s.Text(centerX, 65, "Top Languages", "text-anchor:middle; font:bold 45pt Menlo; fill:white")
	s.Text(centerX, 130, "(By Time Coding)", "text-anchor:middle; font:bold 45pt Menlo; fill:white")
	s.Text(centerX, 170, "In the last year*", "text-anchor:middle; font:bold 18pt Menlo; fill:white")

	for i, language := range languages.Languages {
		s.Rect(950, i*52+200, 60, 30, "fill:"+language.Colour)
		s.Text(1020, i*52+230, language.Name, "font:bold 35pt Menlo; text-anchor:left; fill:"+language.Colour)
		s.Text(1330, i*52+230, fmt.Sprintf("%.1f%%", language.Percent), "font:bold 35pt Menlo; text-anchor:left; fill:"+language.Colour)
	}

	s.Text(centerX, 900, "Graphs by Zac, now in Golang", "text-anchor:middle; font:bold 18pt Menlo; fill:white")
}
