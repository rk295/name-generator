package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	generator "github.com/rk295/name-generator/lib"
)

var (
	rootCmd = &cobra.Command{
		Use:   "name-generator",
		Short: "A random name generator",
		Run:   generate,
	}

	// Options
	number      int
	types       []string
	separator   string
	randomNumer bool
)

func init() {
	allTypes, err := generator.PossibleTypes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().IntVarP(&number, "number", "n", 1, "Number of names to generate")
	rootCmd.PersistentFlags().StringSliceVarP(&types, "types", "t", allTypes, "Types to include")
	rootCmd.PersistentFlags().StringVarP(&separator, "separator", "s", "-", "Separator to use between words")
	rootCmd.PersistentFlags().BoolVarP(&randomNumer, "random", "r", false, "Append a random 6 digit number")
}

// Execute is respomnsible for executing the viper command
func Execute() error {
	return rootCmd.Execute()
}

// generate is the entry point into the name generator from the root Cobra cmd
func generate(cmd *cobra.Command, args []string) {

	if err := generator.CheckType(types); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate the number of names we were asked for
	n := 1
	for n <= number {
		name, err := generator.GetName(types, separator, randomNumer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(name)
		n++
	}
}
