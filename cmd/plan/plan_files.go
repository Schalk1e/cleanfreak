package plan

import (
	"fmt"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
)

type PlanFiles struct {
	dir       string
	move_dirs []string
	to_delete []string
	to_move   map[string][]string
}

func (p PlanFiles) alreadySelected() []string {
	if p.to_move == nil {
		p.to_move = make(map[string][]string)
	}
	if p.to_delete == nil {
		p.to_delete = []string{}
	}

	return append(cmdutil.MapValuesFlatten(p.to_move), p.to_delete...)
}

func (p *PlanFiles) ToDelete() {
	p.to_delete = cmdutil.FileTreeSelect(
		p.dir,
		"Mark files for deletion:",
		p.alreadySelected(),
		false,
	).SelectedFiles
}

func (p *PlanFiles) ToMove() {
	p.to_move = make(map[string][]string)

	for _, d := range p.move_dirs {
		p.to_move[d] = cmdutil.FileTreeSelect(
			p.dir,
			fmt.Sprintf("Mark files to be moved to %s:", d),
			p.alreadySelected(),
			false,
		).SelectedFiles
	}
}

// TODO: Add a plan print section here that outlines the plan.
// Also a method or two that allows us to output this plan as yaml for later.
// Consider using the bubbletea style?
// What about some sort of a printing interface to switch between.
