package wakatime

import (
	"math"

	svg "github.com/ajstarks/svgo"
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

//     ctx.strokeStyle = "white";
// 	ctx.beginPath();
// 	ctx.moveTo(middle.x + ox, middle.y + oy);
// 	ctx.arc(middle.x+ox, middle.y+oy, radius, start+offset*0.0001, end-offset*0.0001);
// 	ctx.lineTo(middle.x+ox, middle.y+oy);
// 	ctx.lineWidth = 5; //offset*0.3;
// 	ctx.stroke();
// 	ctx.fill();
// }

func CreatePiGraph(s *svg.SVG, languages Languages) {
	s.Circle(480, 460, 400, "fill:none;stroke:white;stroke-width:5")
	totalAngle := 0.0
	for _, language := range languages.Languages {
		if language.Percent > 50 {
			language.Percent = 47.5
		}
		if language.Percent > 0.2 {
			CreateSlice(s, totalAngle, language.Percent, language.Colour)
			totalAngle += language.Percent
		}
	}
}
