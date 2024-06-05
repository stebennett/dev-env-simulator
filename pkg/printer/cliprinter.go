package printer

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/table"
	"github.com/stebennett/env-usage-simulator/pkg/simulator"
)

func PrintSimulatorFields(output io.Writer, team *simulator.Team) {
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
