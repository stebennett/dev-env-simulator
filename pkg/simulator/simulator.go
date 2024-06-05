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
	Service Service
	Version int
	Age     int
}

type Environment struct {
	Name     string
	Services []EnvironmentService
}

type Team struct {
	Name           string
	Backlog        []*BacklogItem
	WorkInProgress []*BacklogItem
}

type BacklogItem struct {
	Key     string
	Size    int
	Service *Service
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
			Services: []EnvironmentService{},
		}

		// Build Services
		for _, service := range s.Services {
			service := EnvironmentService{
				Service: *service,
				Version: 1,
				Age:     0,
			}

			env.Services = append(env.Services, service)
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
				Key:     fmt.Sprintf("item-%d", i),
				Size:    pickCycleTime(minCycleTime, maxCycleTime),
				Service: pickService(s.Services),
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
}

func pickService(services []*Service) *Service {
	index := rand.Intn(len(services))
	return services[index]
}

func pickCycleTime(minCycleTime int, maxCycleTime int) int {
	return rand.Intn(maxCycleTime-minCycleTime) + minCycleTime
}
