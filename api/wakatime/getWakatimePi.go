package wakatime

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func GetWakatimePiSVG(w http.ResponseWriter, r *http.Request) {
	languages := ReadLanguages()
	s := svg.New(w)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	CreatePiGraph(s, languages)
	CreateLegend(s, languages)
	s.End()
}
