/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/aoimaru/finop/lib"
	"github.com/spf13/cobra"
)

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
		datas := GetToJ(args[ZERO])
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
	tplCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pplCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pplCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
