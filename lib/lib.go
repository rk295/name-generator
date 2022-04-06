package lib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rk295/name-generator/data"
)

type Permutations map[string][]string

const (
	dataFileSuffix = ".txt"
)

// checkTypes checks slice types against the list of known types. Returns an
// error if the requested type is invalid
func CheckType(types []string) error {
	allTypes := PossibleTypes()
	for _, t := range types {
		if !contains(allTypes, t) {
			return errors.Errorf("type %s is not valid. Possible values are: %s", t, strings.Join(allTypes, ", `"))
		}
	}
	return nil
}

// get an actual name from the list of permuations
func GetName(types []string, separator string, randomNumer bool) (string, error) {

	perms, err := readData(types)
	if err != nil {
		return "", err
	}

	var name []string
	for t := range perms {
		thing := perms[t][ran(len(perms[t]))]
		name = append(name, thing)
	}
	if randomNumer {
		name = append(name, fmt.Sprintf("%d", randomNumber()))
	}
	return strings.Join(name, separator), nil
}

// PossibleTypes returns a string slice of all possible data types. (ls data/*.txt)
func PossibleTypes() []string {
	var names []string
	for _, n := range data.AssetNames() {
		names = append(names, strings.TrimSuffix(n, dataFileSuffix))
	}
	return names
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

// ran picks a random positive int from 0 to max
func ran(max int) int {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return r.Intn(max)
}

func randomNumber() int {
	min := 100000
	max := 999999
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// readData reads in the data for the types requested (eg: colour,dog,etc.)
func readData(types []string) (Permutations, error) {
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
