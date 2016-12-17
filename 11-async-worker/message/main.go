package message

type Message struct {
  command string
  aggregate int
}

func NewCollection(names map[string]int) []Message {
  msgs := []Message{}
  for cmd, aid := range names {
    msgs = append(msgs, Message{command: cmd, aggregate: aid})
  }
  return msgs
}
