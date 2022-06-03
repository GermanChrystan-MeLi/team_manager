package names

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
)

//======================================================================//
type FullName struct {
	FirstName string
	LastName  string
}

//======================================================================//
func CreateNewChooser(choices []wr.Choice) (*wr.Chooser, error) {
	return wr.NewChooser(choices...)
}

//======================================================================//
func CreateFirstName(firstnames []wr.Choice) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	firstNamesChooser, err := CreateNewChooser(firstnames)
	if err != nil {
		return "", err
	}

	firstNameLength := rand.Intn(2) + 1 // 1 or 2
	firstNames := make([]string, 0)

	for i := 0; i < firstNameLength; i++ {
		newName := firstNamesChooser.Pick().(string)
		firstNames = append(firstNames, newName)
	}

	return strings.Join(firstNames, " "), nil
}

//======================================================================//
func CreateLastName(lastnames []wr.Choice) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	lastNameChooser, err := CreateNewChooser(lastnames)
	if err != nil {
		return "", err
	}
	return lastNameChooser.Pick().(string), nil
}

//======================================================================//
func CreateFullName(firstnames, lastnames []wr.Choice) (FullName, error) {
	fn, err := CreateFirstName(firstnames)
	if err != nil {
		return FullName{}, err
	}
	ln, err := CreateLastName(lastnames)
	if err != nil {
		return FullName{}, err
	}
	return FullName{
		FirstName: fn,
		LastName:  ln,
	}, nil
}

//======================================================================//
func CreateFullNameByNationality(nationality string) (FullName, error) {
	switch nationality {
	case "arg":
		return CreateFullName(ArgentinianFirstNames, ArgentinianLastNames)
	case "brz":
		return CreateFullName(BrazilianFirstNames, BrazilianLastNames)
	case "chl":
		return CreateFullName(ChileanFirstNames, ChileanLastNames)
	default:
		return FullName{}, errors.New("nationality is not valid")
	}
}

//======================================================================//
