package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jing/internal/config"
	"jing/internal/tv"
	_ "jing/internal/tv/porn"
	"jing/pkg/util"
	"os"
)

var (
	kind   string
	name   string
	output string
	tvCmd  = &cobra.Command{
		Use:   "tv",
		Short: "search tv",
		Long:  `search tv with kind and name`,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}
)

func init() {

	tvCmd.Flags().StringVarP(&kind, "kind", "k", "porn", "tv kind")
	tvCmd.Flags().StringVarP(&name, "name", "n", "", "web name")
	tvCmd.Flags().StringVarP(&output, "output", "o", "tmp.txt", "output file")

	rootCmd.AddCommand(tvCmd)
}

func run(cmd *cobra.Command, args []string) {
	run0(kind, name)
}

func run0(kind string, name string) {
	fmt.Println("search tv with " + kind + " " + name)
	tvItem := config.ListTv(kind, name)
	fmt.Println(tvItem)
	result := tv.DoSearch(kind, tvItem)

	writeOutput(result)
}

func writeOutput(arr []string) {
	path := util.GetWorkDir() + "/" + output
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range arr {
		f.WriteString(r + "\n")
	}

}
