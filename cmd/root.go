/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/aoimaru/finop/lib"
	"github.com/spf13/cobra"
)

type RootOpt struct {
	jsn string
}

var (
	rOpt = &RootOpt{}
)

type Handle interface {
	get2stc() []lib.ToExt
}

type Target interface {
	getTarget()
}

// type ToExt struct {
// 	Name       string   `json:"name"`
// 	Extentions []string `json:"extentions"`
// }

// func (t2tag TplTag) toOrigin() lib.ToExt {
// 	data := lib.ToExt{
// 		Name:       t2tag.fn,
// 		Extentions: t2tag.tpls,
// 	}
// 	return data
// }

func GetJsonObj(datas *[]lib.ToExt) ([]uint8, error) {
	jsonBlob, err := lib.ToJson(datas)
	if err != nil {
		return nil, err
	}
	return jsonBlob, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finop",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.finop.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&opt.jsn, "json", "j", false, "to json file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
