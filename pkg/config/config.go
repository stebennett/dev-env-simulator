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
	TestingCycleTime     int
}

func ParseConfig() (*Config, error) {
	// Create a new instance of Config
	config := &Config{}

	flag.IntVar(&config.NumberOfServices, "services", 1, "The number of services to simulate")
	flag.IntVar(&config.NumberOfTeams, "teams", 1, "The number of teams to simulate")
	flag.IntVar(&config.NumberOfEnvironments, "environments", 1, "The number of environments to simulate")
	flag.IntVar(&config.NumberOfCycles, "cycles", 20, "The number of cycles to simulate")
	flag.IntVar(&config.MinCycleTime, "minCycleTime", 1, "The minimum cycle time in seconds")
	flag.IntVar(&config.MaxCycleTime, "maxCycleTime", 5, "The maximum cycle time in seconds")
	flag.IntVar(&config.TestingCycleTime, "testingCycleTime", 1, "The testing cycle time in seconds")

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

	if config.TestingCycleTime < 1 {
		return fmt.Errorf("invalid testing cycle time")
	}

	return nil
}
