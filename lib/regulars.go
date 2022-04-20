package lib

import (
	"regexp"
	"strings"
)

func MatchExt(comp string, outs *[]string, extent string) {
	extent = "([a-z]+)." + extent
	r := regexp.MustCompile(extent)
	coms := r.FindAllString(comp, -1)
	for _, com := range coms {
		if len(com) == 0 {
			continue
		}
		*outs = append(*outs, com)
	}
}

func MatchSql(text string) string {
	if strings.Contains(text, "sql = ") {
		words := strings.Split(text, " ")
		var token string
		for _, word := range words {
			if len(word) != 0 {
				token += " " + word
			}
		}
		return token[1:]
	}
	return ""
}
