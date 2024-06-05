package config

import (
	"flag"
	"fmt"
)

type Config struct {
	NumberOfServices     int
	NumberOfTeams        int
	NumberOfEnvironments int
	NumberOfCycles       int
	MinCycleTime         int
	MaxCycleTime         int
	MinTestingCycleTime  int
	MaxTestingCycleTime  int
	TestingWIPLimit      int
	BuildWIPLimit        int
}

func ParseConfig() (*Config, error) {
	// Create a new instance of Config
	config := &Config{}

	flag.IntVar(&config.NumberOfServices, "services", 5, "The number of services to simulate")
	flag.IntVar(&config.NumberOfTeams, "teams", 3, "The number of teams to simulate")
	flag.IntVar(&config.NumberOfEnvironments, "environments", 1, "The number of environments to simulate")
	flag.IntVar(&config.NumberOfCycles, "cycles", 20, "The number of cycles to simulate")
	flag.IntVar(&config.MinCycleTime, "minCycleTime", 1, "The minimum cycle time in seconds")
	flag.IntVar(&config.MaxCycleTime, "maxCycleTime", 5, "The maximum cycle time in seconds")
	flag.IntVar(&config.MinTestingCycleTime, "minTestingCycleTime", 1, "The testing cycle time in seconds")
	flag.IntVar(&config.MaxTestingCycleTime, "maxTestingCycleTime", 5, "The testing cycle time in seconds")
	flag.IntVar(&config.TestingWIPLimit, "testingWIPLimit", 1, "The maximum number of items in testing WIP per team")
	flag.IntVar(&config.BuildWIPLimit, "buildWIPLimit", 1, "The maximum number of items in build WIP per team")

	// Parse command line flags
	flag.Parse()

	err := validateConfig(config)

	return config, err
}

func validateConfig(config *Config) error {
	if config.NumberOfServices < 1 {
		return fmt.Errorf("invalid number of services")
	}

	if config.NumberOfTeams < 1 {
		return fmt.Errorf("invalid number of teams")
	}

	if config.NumberOfEnvironments < 1 {
		return fmt.Errorf("invalid number of environments")
	}

	if config.NumberOfCycles < 1 {
		return fmt.Errorf("invalid number of cycles")
	}

	if config.MinCycleTime < 1 {
		return fmt.Errorf("invalid minimum cycle time")
	}

	if config.MaxCycleTime <= config.MinCycleTime {
		return fmt.Errorf("invalid maximum cycle time")
	}

	if config.MinTestingCycleTime < 1 {
		return fmt.Errorf("invalid testing cycle time")
	}

	if config.MaxTestingCycleTime <= config.MinTestingCycleTime {
		return fmt.Errorf("invalid maximum testing cycle time")
	}

	return nil
}
