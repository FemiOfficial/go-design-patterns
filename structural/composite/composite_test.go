package composite

import "testing"

func TestComposite(t *testing.T) {
	swimmer := AnimalSwimmer {
		Swim: Swim,
	}
	athleteSwimmer := AthleteSwimmer {
		Swim: Swim,
	}	

	swimmer.Eat()
	athleteSwimmer.Train()

	result := swimmer.Swim()
	if result != "swam" {
		t.Errorf("Swim result: %s; want swam", result)
	}

	athleteResult := athleteSwimmer.Swim()
	if athleteResult != "swam" {
		t.Errorf("Swim result: %s; want swam", athleteResult)
	}
}