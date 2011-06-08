package mir

import "os"

func (self *inst) Target() MIRVar {
    if v, ok := self.z.(MIRVar); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Target function."))
}

func (self *inst) X() MIROperand {
    if v, ok := self.x.(MIROperand); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the X function."))
}

func (self *inst) Op() MIROperator {
    if v, ok := self.o.(MIROperator); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Op function."))
}

func (self *inst) Y() MIROperand {
    if v, ok := self.y.(MIROperand); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Y function."))
}

func (self *inst) Cond() MIROperand {
    if v, ok := self.y.(MIROperand); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Cond function."))
}

func (self *inst) Type() MIRType {
    if v, ok := self.y.(MIRType); ok {
        return v
    }
    panic(os.NewError("Instruction does not support the Type function."))
}
