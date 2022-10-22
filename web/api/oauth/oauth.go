package oauth

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/api/wakatime"
	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/db/db_service"
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils/mail"
)

func OAuthHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	code := c.Query("code")
	data, err := GetWakatimeToken(code)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid code")
		return
	}
	userResp, err := GetWakatimeUser(data.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error getting user data")
		return
	}
	user := models.User{
		Username:     userResp.Data.Username,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}

	_, err = db_service.GetUser(user.Username)
	if err == nil {
		c.String(http.StatusAccepted, "User already authorized")
		return
	}

	err = db_service.SaveUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error saving user data")
		return
	}

	err = wakatime.SetCodeData(&user)
	mail.SendMessage(config.CONFIG.ContactEmail, "Zac", "CodeGraphs: New user has joined: "+user.Username, "New user: "+user.Username+" has joined wakatime. Email them here: "+userResp.Data.Email)
	var welcome_message string
	if err != nil {
		welcome_message = `<h1>Welcome to CodeGraphs!</h1>
		<h3>There was an issue with your graphs: ` + err.Error() + `</h3>
		
		<p>Reach out to Zac for account verification</p>
		`
	} else {
		welcome_message = `<h1>Welcome to CodeGraphs!</h1>
		<h3>You can view your graphs <a href="https://graphs.insomnizac.xyz/` + user.Username + `">here</a></h3>
		
		<p>Reach out to Zac for account verification</p>
		`
	}
	mail.SendMessage(userResp.Data.Email, user.Username, "Welcome to CodeGraphs", welcome_message)

	location := url.URL{Path: "/" + user.Username}
	c.Redirect(http.StatusFound, location.RequestURI())
}
