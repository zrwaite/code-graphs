package wakatime

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/zrwaite/github-graphs/models"
)

func CreateLegend(s *svg.SVG, languages models.Languages) {
	centerX := 1200

	s.Text(centerX, 65, "Top Languages", "text-anchor:middle; font:bold 45pt Menlo; fill:white")
	s.Text(centerX, 130, "(By Time Coding)", "text-anchor:middle; font:bold 45pt Menlo; fill:white")

	for i, language := range languages.Languages {
		s.Rect(950, i*52+165, 60, 30, "fill:"+language.Colour)
		s.Text(1020, i*52+195, language.Name, "font:bold 35pt Menlo; text-anchor:left; fill:"+language.Colour)
		s.Text(1330, i*52+195, fmt.Sprintf("%.1f%%", language.Percent), "font:bold 35pt Menlo; text-anchor:left; fill:"+language.Colour)
	}
	length := len(languages.Languages)
	s.Rect(950, length*52+165, 60, 30, "fill:white")
	s.Text(1020, length*52+195, "Other:", "font:bold 35pt Menlo; text-anchor:left; fill:white")
	s.Text(1330, length*52+195, fmt.Sprintf("%.1f%%", languages.Other.Percent), "font:bold 35pt Menlo; text-anchor:left; fill:white")
	offset := length*52 + 240
	// fmt.Println(languages.Other.Languages)
	otherLanguageLines := [][]string{}
	index := 0
	languageLength := 0
	for _, language := range languages.Other.Languages {
		languageLength += len(language.Name)
		if languageLength > 22 {
			index++
			languageLength = len(language.Name)
			if index > 1 {
				break
			}
		}
		if len(otherLanguageLines) <= index {
			otherLanguageLines = append(otherLanguageLines, []string{})
		}
		otherLanguageLines[index] = append(otherLanguageLines[index], language.Name)
	}
	s.Rect(950, offset-30, 520, len(otherLanguageLines)*35, "stroke:white; stroke-width:2; fill:none")
	for i, line := range otherLanguageLines {
		languageList := strings.Join(line, ", ")
		s.Text(960, i*28+offset, languageList, "font:bold 20pt Menlo; text-anchor:left; fill:"+"white")
	}

	s.Text(480, 895, "Graphs by Zac - zrwaite", "text-anchor:middle; font:bold 20pt Menlo; fill:white")
}
