package day12

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Program struct {
	ID    int
	Pipes []*Program
}

func (p *Program) AddPipe(child *Program) {
	p.Pipes = append(p.Pipes, child)
}

func (p *Program) Group() map[int]*Program {
	group := make(map[int]*Program, 0)
	p._group(group)
	//
	//
	return group
}

func (p *Program) _group(g map[int]*Program) {
	if _, ok := g[p.ID]; !ok {
		g[p.ID] = p
		for _, c := range p.Pipes {
			c._group(g)
		}
	}
}

type Village struct {
	Programs []*Program
}

func (v *Village) GetProgram(id int) *Program {
	for _, p := range v.Programs {
		if p.ID == id {
			return p
		}
	}
	panic(fmt.Sprintf("Program not found. ID=%d", id))
}

func (v *Village) Group() []map[int]*Program {
	list := make([]map[int]*Program, 0)
	for _, p := range v.Programs {
		found := false
		for i := range list {
			if _, ok := list[i][p.ID]; ok {
				found = true
			}
		}
		if !found {
			list = append(list, p.Group())
		}
	}
	return list
}

func NewVillage(data []string) *Village {
	//
	//
	village := new(Village)
	var err error
	//
	// Build list of programs
	for _, str := range data {
		p := new(Program)
		p.ID, err = strconv.Atoi(strings.TrimSpace(strings.Split(str, "<->")[0]))
		if err != nil {
			log.Panic(err)
		}
		village.Programs = append(village.Programs, p)
	}
	//
	// Connect programs
	for _, str := range data {
		s := strings.Split(str, "<->")
		programId, err := strconv.Atoi(strings.TrimSpace(s[0]))
		if err != nil {
			log.Panic(err)
		}
		for _, pipeStr := range strings.Split(s[1], ",") {
			pipeProgramId, err := strconv.Atoi(strings.TrimSpace(pipeStr))
			if err != nil {
				log.Panic(err)
			}
			village.GetProgram(programId).AddPipe(village.GetProgram(pipeProgramId))
		}
	}
	return village
}
