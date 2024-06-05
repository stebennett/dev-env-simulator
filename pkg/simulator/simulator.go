package simulator

import (
	"fmt"
	"log"
	"math/rand"
)

type Simulator struct {
	Environments []*Environment
	Services     []*Service
	Teams        []*Team
}

type Service struct {
	Name string
}

type EnvironmentService struct {
	Version             int
	Age                 int
	ItemUnderTest       *BacklogItem
	ItemsWaitingForTest []*BacklogItem
}

type Environment struct {
	Name     string
	Services map[string]*EnvironmentService
}

type Team struct {
	Name           string
	Backlog        []*BacklogItem
	WorkInProgress []*BacklogItem
}

type BacklogItem struct {
	Key         string
	Size        int
	TestingSize int
	DeployTo    string
}

func NewSimulator() *Simulator {
	return &Simulator{}
}

func (s *Simulator) Initialise(
	numberOfServices int,
	numberOfEnvironments int,
	numberOfTeams int,
	minCycleTime int,
	maxCycleTime int,
	minTestingCycleTime int,
	maxTestingCycleTime int,
) {

	s.Services = []*Service{}
	// Build services
	for i := 0; i < numberOfServices; i++ {
		service := Service{
			Name: fmt.Sprintf("service-%d", i),
		}

		s.Services = append(s.Services, &service)
	}

	s.Environments = []*Environment{}
	// Build Environment
	for i := 0; i < numberOfEnvironments; i++ {
		env := Environment{
			Name:     fmt.Sprintf("env-%d", i),
			Services: make(map[string]*EnvironmentService),
		}

		// Build Services
		for _, service := range s.Services {
			envService := EnvironmentService{
				Version:             1,
				Age:                 0,
				ItemUnderTest:       nil,
				ItemsWaitingForTest: []*BacklogItem{},
			}

			env.Services[service.Name] = &envService
		}

		s.Environments = append(s.Environments, &env)
	}

	s.Teams = []*Team{}
	// Build some team backlogs
	for i := 0; i < numberOfTeams; i++ {
		team := Team{
			Name:    fmt.Sprintf("team-%d", i),
			Backlog: []*BacklogItem{},
		}

		for i := 1; i <= 10; i++ {
			backlogItem := BacklogItem{
				Key:         fmt.Sprintf("%s-item-%d", team.Name, i),
				Size:        pickCycleTime(minCycleTime, maxCycleTime),
				TestingSize: pickCycleTime(minTestingCycleTime, maxTestingCycleTime),
				DeployTo:    pickService(s.Services).Name,
			}

			team.Backlog = append(team.Backlog, &backlogItem)
		}

		s.Teams = append(s.Teams, &team)
	}
}

func (s *Simulator) Tick() {

	// move items around backlog
	for _, team := range s.Teams {
		if len(team.Backlog) == 0 {
			log.Println("Team has no backlog items")
			continue
		}

		if (len(team.WorkInProgress) > 0) && (team.WorkInProgress[0].Size > 0) {
			log.Println("Team has items in progress. Working...")
			// Work on the item
			item := team.WorkInProgress[0]
			item.Size--
			if item.Size == 0 {
				// Move the item from work in progress to done
				team.WorkInProgress = team.WorkInProgress[1:]
				s.moveToFirstEnvironment(item)
			}
			continue
		} else if len(team.WorkInProgress) == 0 {
			log.Println("Team has no items in progress. Pulling new ticket...")
			// Move the first item from the backlog to the work in progress
			item := team.Backlog[0]
			team.WorkInProgress = append(team.WorkInProgress, item)
			team.Backlog = team.Backlog[1:]
		}
	}

	// Move items around environments in reverse order
	for envIdx := len(s.Environments) - 1; envIdx >= 0; envIdx-- {
		env := s.Environments[envIdx]
		for thisEnvService, envService := range env.Services {
			if envService.Age == 0 && len(envService.ItemsWaitingForTest) == 0 && envService.ItemUnderTest == nil {
				log.Println("Service is empty. Waiting for testing...")
				continue
			}

			if envService.ItemUnderTest == nil && len(envService.ItemsWaitingForTest) > 0 {
				log.Printf("Service has items waiting for test. Moving to next item...")
				envService.ItemUnderTest = envService.ItemsWaitingForTest[0]
				envService.ItemsWaitingForTest = envService.ItemsWaitingForTest[1:]
				envService.Version++
				envService.Age = 0
			}

			if envService.ItemUnderTest != nil && envService.Age < envService.ItemUnderTest.TestingSize {
				log.Println("Service is in testing. Running tests...")
				envService.Age++
				continue
			}

			if envService.Age == envService.ItemUnderTest.TestingSize {
				log.Println("Service has been tested. Moving to next environment...")

				// move item to next environment
				if envIdx != len(s.Environments)-1 {
					nextEnv := s.Environments[envIdx+1]
					for nextEnvServiceName, nextEnvService := range nextEnv.Services {
						if nextEnvServiceName == thisEnvService {
							nextEnvService.ItemsWaitingForTest = append(nextEnvService.ItemsWaitingForTest, envService.ItemUnderTest)
							break
						}
					}
				}

				envService.Age = 0
				envService.ItemUnderTest = nil

				continue
			}
		}
	}
}

func (s *Simulator) moveToFirstEnvironment(item *BacklogItem) {
	// Move the item to the first environment
	environment := s.Environments[0]
	for serviceName, envService := range environment.Services {
		if serviceName == item.DeployTo {
			envService.ItemsWaitingForTest = append(envService.ItemsWaitingForTest, item)
			return
		}
	}
}

func pickService(services []*Service) *Service {
	index := rand.Intn(len(services))
	return services[index]
}

func pickCycleTime(minCycleTime int, maxCycleTime int) int {
	return rand.Intn(maxCycleTime-minCycleTime) + minCycleTime
}
