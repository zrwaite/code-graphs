package cron

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

/*
import path from 'path'
import fs from 'fs'
import axios from 'axios'
import env from 'dotenv'
env.config()

const swap = (arr: any[], i1: number, i2: number) => {
	let temp = arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = temp
}
const bubbleSortLanguages = (arr: any[]) => {
	for (let i = 0; i < arr.length - 1; i++) {
		for (let j = 0; j < arr.length - i - 1; j++) {
			if (arr[j].time < arr[j + 1].time) swap(arr, j, j + 1)
		}
	}
}

*/

func getCodeData(token string) (models.WakatimeData, error) {
	apiLink := "https://wakatime.com/api/v1/users/current/summaries?timeout=15&writes_only=true"
	// startDate := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	startDate := "2021/10/12"
	date := time.Now().Format("2006-01-02")
	apiLink += "&start=" + startDate + "&end=" + date
	// headers := "{ headers: { Host: \"wakatime.com\", Authorization: \"Bearer " + token + "\" } }"
	resp, err := utils.AuthorizedGetRequest(apiLink, token)
	if err != nil {
		fmt.Println(err)
	}
	var data models.WakatimeData

	if resp.StatusCode == 401 {
		// RefreshAccess(params)
		// return GetRecord(domain, params)
	} else if resp.StatusCode != 200 {
		log.Fatal("Error getting data")
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

/*
	try {
		const codeData: any = await axios.get(apiLink, headers)
		return codeData.data
	} catch (e: any) {
		console.log('Error', e)
		return false
	}
}
*/

func parseLanguages() {
	// ignoreLanguages := []string{"JSON", "Docker", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Bash", "Git Config", "Objective-C", "TOML", "Apache Config", "GitIgnore file", "Shell Script", "GraphQL"}
	// fileName := "json/languages.json"
	// json, err := getCodeData("")
}

/*

const writeLanguages = async () => {
	const folderPath = path.join(__dirname, '../json/')
	let fileName = 'languages.json'
	let filePath = folderPath + fileName
	let json = await getCodeData(`${process.env.WAKATIME_TOKEN}`)
	let languages: any[] = []
	let totalTime: number = 0
	json.data.forEach((data: any) => {
		data.languages.forEach((language: any) => {
			let name = language.name
			let time = language.total_seconds
			let found = false
			for (let i = 0; i < languages.length; i++) {
				if (name != languages[i].name) continue
				languages[i].time += time
				totalTime += time
				found = true
				break
			}
			if (!found) {
				if (!ignoreLanguages.includes(name)) {
					let colour
					switch (name) {
						case 'TypeScript':
							colour = '#0099ff'
							break
						case 'JavaScript':
							colour = '#ecec13'
							break
						case 'C':
							colour = '#666666'
							break
						case 'C#':
							colour = '#9332bf'
							break
						case 'JSON':
							colour = '#339933'
							break
						case 'PHP':
							colour = '#9999ff'
							break
						case 'Python':
							colour = '#0066cc'
							break
						case 'HTML':
							colour = '#ff471a'
							break
						case 'Docker':
							colour = '#1aa3ff'
							break
						case 'SQL':
							colour = '#e6b800'
							break
						case 'Java':
							colour = '#e60000'
							break
						case 'Dart':
							colour = 'rgb(23, 174, 255)'
							break
						case 'SCSS':
							colour = 'rgb(201, 85, 146)'
							break
						case 'CSS':
							colour = 'rgb(28, 49, 220)'
							break
						case 'Rust':
							colour = '#ff5c33'
							break
						case 'Racket':
							colour = 'rgb(100, 13, 20)'
							break
						case 'Markdown':
							colour = '#333333'
							break
						case 'C++':
							colour = 'rgb(83, 136, 200)'
							break
						case 'VHDL':
							colour = 'grey'
							break
						case 'Go':
							colour = 'rgb(20, 156, 206)'
							break
						case 'Swift':
							colour = 'rgb(234, 80, 41)'
							break
						case 'GraphQL':
							colour = 'rgb(215, 0, 135)'
							break
						case 'Svelte':
							colour = 'rgb(235, 62, 39)'
							break
						case 'Ruby':
							colour = 'rgb(217, 10, 0)'
							break
						default:
							colour = 'white'
							break
					}
					languages.push({ name: name, colour: colour, time: time, percent: 0 })
					totalTime += time
				}
			}
			found = false
		})
	})
	bubbleSortLanguages(languages)
	for (let i = 0; i < languages.length; i++) {
		languages[i].percent = Math.round((languages[i].time / totalTime) * 1000) / 10
		languages[i].time = Math.round(languages[i].time)
		if (languages[i].percent <= 0.1) languages.splice(i, languages.length - 1)
	}
	totalTime = Math.round(totalTime)
	if (languages.length > 13) languages = languages.slice(0, 13)
	fs.writeFileSync(filePath, JSON.stringify({ time: totalTime, languages: languages }))
}

export default writeLanguages
*/
