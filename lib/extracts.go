package lib

import (
	"bufio"
	"os"
)

func GetExts(fileName string, extentName string) []string {
	fp, err := os.Open(fileName)
	origins := make([]string, 10, 10)
	if err != nil {
		return nil
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		MatchExt(scanner.Text(), &origins, extentName)
	}
	if err = scanner.Err(); err != nil {
		return nil
	}
	ret := func() []string {
		coms := make([]string, 0)
		for _, origin := range origins {
			if len(origin) != 0 {
				coms = append(coms, origin)
			}
		}
		return coms
	}
	return ret()
}
