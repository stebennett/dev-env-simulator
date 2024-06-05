package main

import (
	"os"

	"github.com/stebennett/env-usage-simulator/pkg/config"
	"github.com/stebennett/env-usage-simulator/pkg/printer"
	"github.com/stebennett/env-usage-simulator/pkg/simulator"
)

func main() {
	conf, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	simulator := simulator.NewSimulator()
	simulator.Initialise(
		conf.NumberOfServices,
		conf.NumberOfEnvironments,
		conf.NumberOfTeams,
		conf.MinCycleTime,
		conf.MaxCycleTime,
		conf.MinTestingCycleTime,
		conf.MaxTestingCycleTime,
	)

	for i := 0; i < conf.NumberOfCycles; i++ {
		simulator.Tick()
		for _, team := range simulator.Teams {
			printer.PrintSimulatorTeam(os.Stdout, team)
		}
		for _, env := range simulator.Environments {
			printer.PrintSimulatorEnvironment(os.Stdout, env)
		}
	}
}
