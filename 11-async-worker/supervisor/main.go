package supervisor

import "sync"

type Message struct {
  command string
  aggregate int
}

func NewMessageCollection(names map[string]int) []Message {
  msgs := []Message{}
  for cmd, aid := range names {
    msgs = append(msgs, Message{command: cmd, aggregate: aid})
  }
  return msgs
}

type Worker func (Message)

type Supervisor struct {
  workers []Worker
  wg sync.WaitGroup
}

type ISupervisor interface {
  CreateWorker(func (Message))
  Handle(Message)
}

func New() *Supervisor {
  return &Supervisor{workers: []Worker{}}
}

func (s *Supervisor) CreateWorker(worker func (Message)) {
  s.workers = append(s.workers, worker)
}

func (s *Supervisor) Wait() {
  s.wg.Wait()
}

func (s *Supervisor) Handle(msg Message) {
  s.wg.Add(1)
  for _, worker := range s.workers {
    go func() {
      defer s.wg.Done()
      worker(msg)
    }()
  }
}
