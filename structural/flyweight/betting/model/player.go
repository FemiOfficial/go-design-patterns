package model

type Player struct {
	ID uint64
	Name string
	PreviousTeams uint64
	Photo []byte
}