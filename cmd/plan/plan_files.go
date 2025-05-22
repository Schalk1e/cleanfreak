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
	FilesToDelete []string
	FilesToMove   map[string][]string
	dir           string
	move_dirs     []string
}

func (p *PlanFiles) ToDelete() {
	p.FilesToDelete = cmdutil.FileTreeSelect(
		p.dir,
		"Mark files for deletion:",
		p.alreadySelected(),
		false,
	).SelectedFiles
}

func (p *PlanFiles) ToMove() {
	p.FilesToMove = make(map[string][]string)

	for _, d := range p.move_dirs {
		p.FilesToMove[d] = cmdutil.FileTreeSelect(
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

func (p PlanFiles) PrintPlan() {
	fmt.Printf(
		"\n You selected: %s file(s) to delete and %s file(s) to move.\n\n",
		p.deleteCount(),
		p.moveCount(),
	)
}

func (p PlanFiles) alreadySelected() []string {
	if p.FilesToMove == nil {
		p.FilesToMove = make(map[string][]string)
	}
	if p.FilesToDelete == nil {
		p.FilesToDelete = []string{}
	}

	return append(cmdutil.MapValuesFlatten(p.FilesToMove), p.FilesToDelete...)
}

func (p PlanFiles) deleteCount() string {
	return strconv.Itoa(len(p.FilesToDelete))
}

func (p PlanFiles) moveCount() string {
	var move_count int

	for _, files := range p.FilesToMove {
		move_count += len(files)
	}

	return strconv.Itoa(move_count)
}
