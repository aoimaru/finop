/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/aoimaru/finop/lib"
	"github.com/spf13/cobra"
)

type Phs struct {
	fps []string
}

type PhsTag struct {
	fn   string
	sqls []string
}

func (p2tag *PhsTag) getTarget() {
	fp, err := os.Open(p2tag.fn)
	origins := make([]string, 10, 10)
	if err != nil {
		log.Fatal(err)
		p2tag.sqls = origins
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		origin := lib.MatchSql(scanner.Text())
		if len(origin) != 0 {
			origins = append(origins, origin)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
		p2tag.sqls = origins
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
	p2tag.sqls = ret()
}

func (p2t Phs) get2stc() []lib.ToExt {
	datas := make([]lib.ToExt, 0)
	for _, fp := range p2t.fps {
		p2tag := new(PhsTag)
		p2tag.fn = fp

		p2tag.getTarget()
		data := lib.ToOrigin(p2tag.fn, p2tag.sqls)
		datas = append(datas, data)
	}
	return datas
}

// pplCmd represents the ppl command
var pplCmd = &cobra.Command{
	Use:   "ppl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		const ZERO = 0
		fps, err := lib.ListFiles(args[ZERO])
		if err != nil {
			log.Fatal(err)
		}
		php2obj := new(Phs)
		php2obj.fps = fps

		var src2hdl Handle = php2obj
		datas := src2hdl.get2stc()

		switch opt.jsn {
		case true:
			lib.ToFile(args[ZERO], datas)
		case false:
			jsonObj, err := GetJsonObj(&datas)
			if err != nil {
				log.Fatal(err)
			}
			lib.ToClean(jsonObj)
		}
	},
}

func init() {
	rootCmd.AddCommand(pplCmd)
	pplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")
	// tplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pplCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pplCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
