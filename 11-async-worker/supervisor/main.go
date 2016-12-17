package supervisor

import "../message"

type Worker struct {
  work func (message.Message, *Worker)
}

type Supervisor struct {
  workers []*Worker
}

type ISupervisor interface {
  CreateWorker(func (message.Message, *Worker))
  Handle(message.Message)
}

func New() Supervisor {
  return Supervisor{workers: []*Worker{}}
}

func (s *Supervisor) CreateWorker(work func (message.Message, *Worker)) {
  s.workers = append(s.workers, &Worker{
    work: work,
  })
}

func (s *Supervisor) Handle(msg message.Message) {
  for _, worker := range s.workers {
    go worker.work(msg, worker)
  }
}
