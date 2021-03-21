package main

import "github.com/dadikovi/race-timer/server/core"

func (a *App) RefreshRace() error {
	race, err := core.GetRaceInstance(a.DB)
	a.race = race
	return err
}
