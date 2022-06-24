package names

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"
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
func CreateFullNameByNationality(country constants.Country) (FullName, error) {
	switch country {
	case constants.Argentina:
		return CreateFullName(ArgentinianFirstNames, ArgentinianLastNames)
	case constants.Brazil:
		return CreateFullName(BrazilianFirstNames, BrazilianLastNames)
	case constants.Chile:
		return CreateFullName(ChileanFirstNames, ChileanLastNames)
	default:
		return FullName{}, errors.New("nationality is not valid")
	}
}

//======================================================================//
