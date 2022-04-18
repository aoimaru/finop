/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

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

func getToJ(directory string) []lib.ToJ {
	datas := make([]lib.ToJ, 0)
	fps, err := lib.ListFiles(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, fp := range fps {
		exts := lib.GetExts(fp, opt.ext)
		data := lib.ToOrigin(fp, exts)
		datas = append(datas, data)
	}
	return datas
}

func getJsonObj(datas *[]lib.ToJ) ([]uint8, error) {
	jsonBlob, err := lib.ToJson(datas)
	if err != nil {
		return nil, err
	}
	return jsonBlob, nil
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
		datas := getToJ(args[ZERO])
		switch opt.jsn {
		case true:
			lib.ToFile(args[ZERO], datas)
		case false:
			jsonObj, err := getJsonObj(&datas)
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
	tplCmd.Flags().StringVarP(&opt.ext, "ext", "e", "php", "get phpFile")
	tplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tplCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tplCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
