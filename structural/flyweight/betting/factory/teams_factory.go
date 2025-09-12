package factory

import (
	"fmt"
	"godesignpatterns/structural/flyweight/betting/model"
)

const (
	TEAM_A = 1
	TEAM_B = 2
)

func GetTeamFactory(team uint64) (model.Team, error) {
	switch team {
	case TEAM_A:
		return model.Team{
			ID: 1,
			Name: "TEAM_A",
			HistoricalData: []model.HistoricalData{
				{Year: 2024, Matches: []model.Match{}},
			},
		}, nil
	case TEAM_B:
		return model.Team{
			ID: 2,
			Name: "TEAM_B",
			HistoricalData: []model.HistoricalData{
				{Year: 2024, Matches: []model.Match{}},
			},
		}, nil
	default:
		return model.Team{}, fmt.Errorf("invalid team")
	}
}


