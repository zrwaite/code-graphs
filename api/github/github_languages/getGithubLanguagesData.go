package github_languages

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetGithubLanguagesData(username string, hide string) (string, error) {
	// https: //github-readme-stats.vercel.app/api/top-langs/?username=zrwaite&hide=makefile,powershell,html,css&layout=compact&langs_count=10&theme=dark
	resp, err := http.Get("https://github-readme-stats.vercel.app/api/top-langs/?username=" + username + "&hide=" + hide + "&layout=compact&langs_count=10&theme=dark")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("github readme stats API returned non-200 status code")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
