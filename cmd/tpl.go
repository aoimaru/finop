/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aoimaru/finop/lib"
	"github.com/spf13/cobra"
)

type Options struct {
	ext string
	jsn bool
}

var (
	opt = &Options{}
)

type Tpl struct {
	fps []string
}

type TplTag struct {
	fn   string
	tpls []string
	ext  string
}

func (t2tag *TplTag) getTarget() {
	fp, err := os.Open(t2tag.fn)
	origins := make([]string, 10, 10)
	if err != nil {
		log.Fatal(err)
		t2tag.tpls = origins
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		lib.MatchExt(scanner.Text(), &origins, t2tag.ext)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
		t2tag.tpls = origins
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
	t2tag.tpls = ret()
}

func (t2t Tpl) get2stc() []lib.ToExt {
	datas := make([]lib.ToExt, 0)
	for _, fp := range t2t.fps {

		t2tag := new(TplTag)
		t2tag.fn = fp
		t2tag.ext = opt.ext

		t2tag.getTarget()
		data := lib.ToOrigin(t2tag.fn, t2tag.tpls)

		datas = append(datas, data)
	}
	return datas
}

// tplCmd represents the tpl command
var tplCmd = &cobra.Command{
	Use:   "tpl",
	Short: "get",
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
		tpl2obj := new(Tpl)
		tpl2obj.fps = fps

		var src2hdl Handle = tpl2obj
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

func listAll(arg string) {
	fmt.Println("hello")
}

func ToName(name string) string {
	return name
}

func init() {
	rootCmd.AddCommand(tplCmd)
	tplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")
	tplCmd.Flags().StringVarP(&opt.ext, "ext", "e", "php", "get phpFile")
	// tplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tplCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tplCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
