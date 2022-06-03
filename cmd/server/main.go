package main

import (
	"fmt"

	"github.com/GermanChrystan-MeLi/team_manager/utils/constants"
	"github.com/GermanChrystan-MeLi/team_manager/utils/names"
)

func main() {
	name, _ := names.CreateFullNameByNationality(constants.Brazil)
	fmt.Println(name)
}
