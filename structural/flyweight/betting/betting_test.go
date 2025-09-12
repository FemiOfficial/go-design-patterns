package betting

import (
	"testing"
	"godesignpatterns/structural/flyweight/betting/model"
)

func TestTeamsFactory(t *testing.T) {
	teamsFactory := &TeamsFactory{
		Teams: make(map[uint64]*model.Team),
	}

	teamA := teamsFactory.GetTeam(1)
	teamB := teamsFactory.GetTeam(2)
	teamC := teamsFactory.GetTeam(3)

	if teamC != nil {
		t.Errorf("Team C should be nil, since no team with id 3 exists")
	}

	if teamA.Name != "TEAM_A" {
		t.Errorf("Team A name is not TEAM_A")
	}

	if teamB.Name != "TEAM_B" {
		t.Errorf("Team B name is not TEAM_B")
	}

	if teamsFactory.GetNumberOfObjects() != 2 {
		t.Errorf("Number of objects is not 2")
	}
}

func TestTeamsFactory_GetTeam_HighVolume(t *testing.T) {
	teamsFactory := &TeamsFactory{
		Teams: make(map[uint64]*model.Team),
	}

	for i := 0; i < 10000; i++ {
		teamsFactory.GetTeam(1)
	}

	for i := 0; i < 10000; i++ {
		teamsFactory.GetTeam(2)
	}

	if teamsFactory.GetNumberOfObjects() != 2 {
		t.Errorf("Number of objects should be 2")
	}
}