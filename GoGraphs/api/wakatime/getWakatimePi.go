package wakatime

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func GetWakatimePiSVG(w http.ResponseWriter, r *http.Request) {
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}
