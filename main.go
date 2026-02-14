package main

import (
	"fmt"
	"maps"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	model := model{
		selected:  make(map[int]struct{}),
		cursor:    0,
		collected: make(map[piece]struct{}),
		needed:    make(map[piece]struct{}),
	}
	for _, p := range allPieces {
		model.needed[p] = struct{}{}
	}

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// flag for whether to recompute collected and needed pieces
	recompute := false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// quit
			return m, tea.Quit

		case "up", "k":
			// move down up
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// move cursor down
			if m.cursor < len(allRooms)-1 {
				m.cursor++
			}

		case "enter", " ":
			// flag that we need to recompute collected and needed pieces
			recompute = true
			// have we already selected this room
			_, ok := m.selected[m.cursor]
			if ok {
				// yes, unselect it. first, delete it from that set
				delete(m.selected, m.cursor)
			} else {
				// no, select the room
				m.selected[m.cursor] = struct{}{}
			}

		}
	}

	if recompute {
		// initialize
		collected := make(map[piece]struct{})
		needed := make(map[piece]struct{})
		// walk selected rooms and flag their pieces as collected
		for i := range m.selected {
			p := allRooms[i].piece
			collected[p] = struct{}{}
		}

		// walk all pieces and mark the ones not collected as needed
		for _, p := range allPieces {
			if _, ok := collected[p]; !ok {
				needed[p] = struct{}{}
			}
		}

		m.collected = collected
		m.needed = needed
	}

	return m, nil
}

func (m model) View() string {
	out := ""

	for i, room := range allRooms {
		// is the cursor pointing at this?
		cursor := ' '
		if m.cursor == i {
			cursor = '>'
		}

		// has it been collected?
		selected := ' '
		if _, ok := m.selected[i]; ok {
			selected = '✔'
		}

		out += fmt.Sprintf("%c %c %s (%s)\n", cursor, selected, room.name, room.piece)
	}
	out += "\nCollected: "
	collected := slices.Collect(maps.Keys(m.collected))
	slices.Sort(collected)
	for _, p := range collected {
		out += fmt.Sprintf("%s ", p)
	}

	needed := slices.Collect(maps.Keys(m.needed))
	slices.Sort(needed)
	out += "\nNeeded: "
	for _, p := range needed {
		out += fmt.Sprintf("%s ", p)
	}
	out += "\n\nctrl-c or q to quit, up/j and down/k to scroll, enter/space to select\n"

	return out
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error running program: %v", err)
		os.Exit(1)
	}
}
