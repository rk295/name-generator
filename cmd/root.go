package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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
	number int
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&number, "number", "n", 1, "Number of names to generate")
}

// Execute is respomnsible for executing the viper command
func Execute() error {
	return rootCmd.Execute()
}

func generate(cmd *cobra.Command, args []string) {
	perms, err := readData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n := 1
	for n <= number {
		fmt.Println(getName(perms))
		n++
	}

}

func getName(perms permutations) string {
	var name []string

	for t := range perms {
		thing := perms[t][ran(len(perms[t]))]
		name = append(name, thing)

	}
	return fmt.Sprintf(strings.Join(name, "-"))

}

// ran picks a random positive int from 0 to max
func ran(max int) int {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return r.Intn(max)
}

// readData pulls all the data out of the files and returns the permutation
func readData() (permutations, error) {

	perms := make(map[string][]string)

	for _, asset := range data.AssetNames() {
		_, file := filepath.Split(asset)
		fileName := file[0 : len(file)-4] // Hacky, rip off .txt

		data, err := readLines(asset)
		if err != nil {
			return perms, err
		}
		perms[fileName] = data

	}
	// perms = allData
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
