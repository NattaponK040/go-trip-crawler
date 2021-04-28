package process

import (
	"go-crawler/process/pkg"
	"go-crawler/repository"
)

type Processor struct {
	trip *pkg.TripadvisorProc
}

func NewProcessor(repo *repository.MgRepository) *Processor {
	return &Processor{
		trip: pkg.NewTripasdvisor(repo),
	}
}

func (p *Processor) Runproc()  {
	p.trip.Runprocess()
}

