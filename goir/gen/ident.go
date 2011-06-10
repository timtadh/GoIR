package gen

import "fmt"

type Ident struct {
    name   string
    public bool
}

func (self *Ident) Name() string { return self.name }
func (self *Ident) Public() bool { return self.public }

func (self *Ident) String() string {
    return fmt.Sprintf("<Ident %v %v>", self.name, self.public)
}
