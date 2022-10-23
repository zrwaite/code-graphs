package wakatime_top_stats

import (
	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/utils"
)

func CreateList(s *svg.SVG, title string, list []string, x_middle int, x_left int, y int, colour bool) {
	s.Text(x_middle, y, title, "text-anchor:middle;font:bold 42pt Menlo;fill:white")
	y += 80
	for _, item := range list {
		colourString := "white"
		if colour {
			colourString = utils.GetColour(item)
		}
		s.Text(x_left, y, item, "font:35pt Menlo;fill:"+colourString)
		y += 60
	}
}
