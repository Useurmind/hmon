package main

type Closable interface {
	Close()
}

type Closer struct {
	closables []Closable
}

func NewCloser() *Closer {
	return &Closer{
		closables: make([]Closable, 0),
	}
}

func (c *Closer) AddClosable(closable Closable) {
	c.closables = append(c.closables, closable)
}

func (c *Closer) Close() {
	for _, closable := range c.closables {
		closable.Close()
	}
}
