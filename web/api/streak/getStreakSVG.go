package streak

import "github.com/gin-gonic/gin"

func GetStyleSheet() string {
	return `<style>
	.ContributionCalendar-day {outline: rgba(255, 255, 255, 0.05);}
	.ContributionCalendar-day[data-level="1"] {fill: #0e4429;}
	.ContributionCalendar-day[data-level="2"] {fill: #006d32;}
	.ContributionCalendar-day[data-level="3"] {fill: #26a641;}
	.ContributionCalendar-day[data-level="4"] {fill: #39d353;}
	.ContributionCalendar-label {
		font-size: 12px;
		fill: #c9d1d9;
	}
	/*
	svg {
		padding: 2px;
		background-color: black;
	}
	*/
	svg > g {
		fill: black;
	}
	</style>
	`
}

func GetBackground() string {
	return `
	<rect width="843" height="148" x="0" y="0" fill="black" rx="10" stroke="white"/>
	`
}

func GetStreakSVG(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.String(400, "Username is required")
		return
	}

	html, err := GetStreakData(username)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	svg := GetSvg(html)
	streakGraphBytes := []byte(svg)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "public, max-age=3600")
	c.Data(200, "image/svg+xml", streakGraphBytes)
}
