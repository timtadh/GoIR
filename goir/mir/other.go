package mir

import "os"
import "goir/mir/kinds"

func (self *inst) Kind() kinds.Kind {
    return self.kind
}

func (self *inst) Label() MIRLabel {
    if v, ok := self.z.(MIRLabel); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Label function."))
}

func (self *inst) Param() MIRVar {
    if v, ok := self.z.(MIRVar); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Param function."))
}

func (self *inst) Pos() uint {
    if v, ok := self.x.(uint); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Pos function."))
}
