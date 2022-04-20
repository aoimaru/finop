/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
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

func GetToJ(directory string) []lib.ToJ {
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

func GetJsonObj(datas *[]lib.ToJ) ([]uint8, error) {
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
	rootCmd.Flags().StringVarP(&rOpt.jsn, "json", "j", "default.json", "to json file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
