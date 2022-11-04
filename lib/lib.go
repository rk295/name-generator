package lib

import (
	"embed"
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"
)

type Permutations map[string][]string

const (
	dataDirName    = "data"
	dataFileSuffix = ".txt"

	// If the use wants a random number appended to the name these control the
	// min and max of that number, i.e. it will be within this range
	randonmNumberMin = 100000
	randonmNumberMax = 999999
)

//go:embed data/*.txt
var nameList embed.FS

// checkTypes checks slice types against the list of known types. Returns an
// error if the requested type is invalid
func CheckType(types []string) error {
	allTypes, err := PossibleTypes()
	if err != nil {
		return err
	}
	for _, t := range types {
		if !contains(allTypes, t) {
			return fmt.Errorf("type %s is not valid. Possible values are: %s", t, strings.Join(allTypes, ", `"))
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
	for _, t := range types {
		thing := perms[t][ran(len(perms[t]))]
		name = append(name, thing)
	}
	if randomNumer {
		name = append(name, fmt.Sprintf("%d", randomNumber()))
	}
	return strings.Join(name, separator), nil
}

// PossibleTypes returns a string slice of all possible data types. (ls data/*.txt)
func PossibleTypes() ([]string, error) {
	var names []string

	files, err := nameList.ReadDir(dataDirName)
	if err != nil {
		return names, err
	}

	for _, n := range files {
		names = append(names, strings.TrimSuffix(n.Name(), dataFileSuffix))
	}
	return names, nil
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

// Initialise the rand package
func newRand() *rand.Rand {
	rand.Seed(time.Now().UnixNano())
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

// ran picks a random positive int from 0 to max
func ran(max int) int {
	r := newRand()
	return r.Intn(max)
}

// randomNumber returns a random 6 digit int
func randomNumber() int {
	r := newRand()
	return r.Intn(randonmNumberMax-randonmNumberMin) + randonmNumberMin
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
func readLines(name string) ([]string, error) {
	data, err := nameList.ReadFile(path.Join(dataDirName, name))
	if err != nil {
		return []string{}, err
	}
	return strings.Fields(string(data)), err
}
