package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/db/db_service"
	"github.com/zrwaite/github-graphs/models"
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

	err = db_service.SaveUser(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println(user)
		c.String(http.StatusInternalServerError, "Internal server error saving user data")
		return
	}

	c.String(http.StatusOK, "Success")
}
