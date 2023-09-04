package utils

import "github.com/zrwaite/github-graphs/models"

type GTLTType interface {
	int | string
}

type SortTrait interface {
	models.Language | GTLTType
}

func GenericSort[V SortTrait](list []V, lt func(a, b V) bool) []V {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if lt(list[j], list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func SortLanguagesByTime(languages []models.Language) []models.Language {
	return GenericSort(languages, func(a, b models.Language) bool {
		return a.TotalSeconds < b.TotalSeconds
	})
}
