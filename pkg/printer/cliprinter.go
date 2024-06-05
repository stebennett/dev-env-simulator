package printer

import (
	"fmt"
	"io"
	"sort"

	"github.com/jedib0t/go-pretty/table"
	"github.com/stebennett/env-usage-simulator/pkg/simulator"
)

func PrintSimulatorTeam(output io.Writer, team *simulator.Team) {
	t := table.NewWriter()
	t.SetOutputMirror(output)
	t.SetTitle(fmt.Sprintf("Team %s", team.Name))

	t.AppendHeader(table.Row{"State", "Tickets"})

	ticketNames := []string{}
	for _, ticket := range team.Backlog {
		s := fmt.Sprintf("%s (%d)", ticket.Key, ticket.Size)
		ticketNames = append(ticketNames, s)
	}

	t.AppendRow([]interface{}{"To Do", ticketNames})

	ticketNames = []string{}
	for _, ticket := range team.WorkInProgress {
		s := fmt.Sprintf("%s (%d)", ticket.Key, ticket.Size)
		ticketNames = append(ticketNames, s)
	}

	t.AppendRow([]interface{}{"In Progress", ticketNames})

	t.Render()
}

func PrintSimulatorEnvironment(output io.Writer, env *simulator.Environment) {
	t := table.NewWriter()
	t.SetOutputMirror(output)
	t.SetTitle(fmt.Sprintf("Environment %s", env.Name))

	t.AppendHeader(table.Row{"Service", "Version", "Age", "Item Under Test", "Items Waiting For Test"})

	services := make([]string, 0, len(env.Services))
	for k, _ := range env.Services {
		services = append(services, k)
	}
	sort.Strings(services)

	for _, serviceName := range services {
		service := env.Services[serviceName]

		itemUnderTest := ""
		if service.ItemUnderTest != nil {
			itemUnderTest = fmt.Sprintf("%s (%d)", service.ItemUnderTest.Key, service.ItemUnderTest.TestingSize)
		}

		itemsWaitingForTest := []string{}
		for _, item := range service.ItemsWaitingForTest {
			itemsWaitingForTest = append(itemsWaitingForTest, item.Key)
		}

		t.AppendRow([]interface{}{
			serviceName,
			service.Version,
			service.Age,
			itemUnderTest,
			itemsWaitingForTest,
		})
	}

	t.Render()
}
