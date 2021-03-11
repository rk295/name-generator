package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/rk295/name-generator/data"
)

type permutations map[string][]string

var (
	rootCmd = &cobra.Command{
		Use:   "name-generator",
		Short: "A random name generator",
		Run:   generate,
	}

	// Options
	number    int
	types     []string
	separator string
)

const (
	dataFileSuffix = ".txt"
)

func init() {
	allTypes := possibleTypes()

	rootCmd.PersistentFlags().IntVarP(&number, "number", "n", 1, "Number of names to generate")
	rootCmd.PersistentFlags().StringSliceVarP(&types, "types", "t", allTypes, "Types to include")
	rootCmd.PersistentFlags().StringVarP(&separator, "separator", "s", "-", "Separator to use between words")
}

func possibleTypes() []string {
	var names []string
	for _, n := range data.AssetNames() {
		names = append(names, strings.TrimSuffix(n, dataFileSuffix))
	}
	return names
}

// Execute is respomnsible for executing the viper command
func Execute() error {
	return rootCmd.Execute()
}

// generate is the entry point into the name generator from the root Cobra cmd
func generate(cmd *cobra.Command, args []string) {
	perms, err := readData(types)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate the number of names we were asked for
	n := 1
	for n <= number {
		fmt.Println(getName(perms))
		n++
	}
}

// get an actual name from the list of permuations
func getName(perms permutations) string {
	var name []string
	for t := range perms {
		thing := perms[t][ran(len(perms[t]))]
		name = append(name, thing)
	}
	return fmt.Sprintf(strings.Join(name, separator))
}

// ran picks a random positive int from 0 to max
func ran(max int) int {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return r.Intn(max)
}

// readData reads in the data for the types requested (eg: colour,dog,etc.)
func readData(types []string) (permutations, error) {
	perms := make(map[string][]string)
	for _, asset := range types {

		data, err := readLines(fmt.Sprintf("%s%s", asset, dataFileSuffix))
		if err != nil {
			return perms, err
		}
		perms[asset] = data
	}
	return perms, nil
}

// readLines returns the string slice of the specified file in data/
func readLines(path string) ([]string, error) {
	data, err := data.Asset(path)
	if err != nil {
		return []string{}, err
	}
	return strings.Fields(string(data)), err
}
