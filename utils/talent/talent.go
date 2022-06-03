package talent

import (
	"errors"
	"math/rand"
	"time"

	"github.com/GermanChrystan-MeLi/team_manager/utils/constants"
	wr "github.com/mroth/weightedrand"
)

var TalentChoices = []wr.Choice{
	{Item: constants.Low, Weight: 1},
	{Item: constants.Common, Weight: 2},
	{Item: constants.Good, Weight: 3},
	{Item: constants.Great, Weight: 3},
	{Item: constants.Excellent, Weight: 2},
	{Item: constants.Messi, Weight: 1},
}

func GetRandomTalent() (constants.Talent, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	talentChooser, err := wr.NewChooser(TalentChoices...)
	if err != nil {
		return 0, errors.New("could not create random talent stat")
	}
	return talentChooser.Pick().(constants.Talent), nil
}
