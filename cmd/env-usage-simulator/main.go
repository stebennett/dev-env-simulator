package main

import (
	"log"
	"os"

	"github.com/stebennett/dev-env-simulator/pkg/config"
	"github.com/stebennett/dev-env-simulator/pkg/printer"
	"github.com/stebennett/dev-env-simulator/pkg/simulator"
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
		conf.BuildWIPLimit,
		conf.TestingWIPLimit,
	)

	for i := 0; i < conf.NumberOfCycles; i++ {
		log.Printf("=========== CYCLE %d ==========", i)
		simulator.Tick()
		for _, team := range simulator.Teams {
			printer.PrintSimulatorTeam(os.Stdout, team)
		}
		for _, env := range simulator.Environments {
			printer.PrintSimulatorEnvironment(os.Stdout, env)
		}
	}
}
