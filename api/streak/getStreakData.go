package streak

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetStreakData(username string) (string, error) {
	resp, err := http.Get("https://github.com/users/" + username + "/contributions")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("github API returned non-200 status code")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
