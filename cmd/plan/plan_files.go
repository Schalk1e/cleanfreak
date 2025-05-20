package plan

import (
	"fmt"
	"log"
	"os"
	"strconv"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"gopkg.in/yaml.v2"
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

func (p PlanFiles) OutputPlan(path string) {
	plan, err := yaml.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, plan, 0666); err != nil {
		log.Fatal(err)
	}
}

func (p PlanFiles) deleteCount() string {
	return strconv.Itoa(len(p.to_delete))
}

func (p PlanFiles) moveCount() string {
	var move_count int

	for _, files := range p.to_move {
		move_count += len(files)
	}

	return strconv.Itoa(move_count)
}

func (p PlanFiles) PrintPlan() {
	fmt.Printf(
		"\n You selected: %s file(s) to delete and %s file(s) to move.\n\n",
		p.deleteCount(),
		p.moveCount(),
	)
}
