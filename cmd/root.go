package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
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
	number      int
	types       []string
	separator   string
	randomNumer bool
)

const (
	dataFileSuffix = ".txt"
)

func init() {
	allTypes := possibleTypes()

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

	err := checkType(types)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	if randomNumer {
		name = append(name, fmt.Sprintf("%d", randomNumber()))
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

func randomNumber() int {
	min := 100000
	max := 999999
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// containers looks for string e in slice s. Returns true if found, false if not
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// checkTypes checks slice types against the list of known types. Returns an
// error if the requested type is invalid
func checkType(types []string) error {
	allTypes := possibleTypes()
	for _, t := range types {
		if !contains(allTypes, t) {
			return errors.Errorf("type %s is not valid. Possible values are: %s", t, strings.Join(allTypes, ", `"))
		}
	}
	return nil
}

// possibleTypes returns a string slice of all possible data types. (ls data/*.txt)
func possibleTypes() []string {
	var names []string
	for _, n := range data.AssetNames() {
		names = append(names, strings.TrimSuffix(n, dataFileSuffix))
	}
	return names
}
