package betting

import (
	"godesignpatterns/structural/flyweight/betting/model"
	factory "godesignpatterns/structural/flyweight/betting/factory"
)

type TeamsFactory struct {
	Teams map[uint64]*model.Team
}

func (t *TeamsFactory) GetTeam(id uint64) *model.Team {
	existingTeam := t.Teams[id]
	if existingTeam != nil {
		return existingTeam
	}

	team, err := factory.GetTeamFactory(id)
	if err != nil {
		return nil
	}

	t.Teams[id] = &team
	return 	t.Teams[id]	
}

func (t *TeamsFactory) GetNumberOfObjects() int {
	return len(t.Teams)
}