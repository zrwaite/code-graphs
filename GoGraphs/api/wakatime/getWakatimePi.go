package wakatime

import (
	"fmt"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func GetWakatimePiSVG(w http.ResponseWriter, r *http.Request) {
	languages := ReadLanguages()
	fmt.Println(languages.TotalTime)
	s := svg.New(w)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	// s.Circle(250, 250, 125, "fill:none;stroke:white")
	CreatePiGraph(s, languages)
	s.End()
}
