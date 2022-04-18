package lib

import "regexp"

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
