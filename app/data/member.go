package data

import (
	"log"
	"os-container-project/model"
	"os-container-project/pkg/util"
)

func LoadMembersData() []model.Member {
	var members []model.Member

	if err := util.LoadJSONFile("data/members.json", &members); err != nil {
		log.Fatal("Error loading members data: ", err)
	}

	return members
}
