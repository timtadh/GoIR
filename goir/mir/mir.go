package mir

import "goir/mir/kinds"

type inst struct {
    kind kinds.Kind
    o    MIROperator
    x    MIROperand
    y    MIROperand
    z    MIROperand
}

type MIRInst interface {
    Kind() kinds.Kind
}

type Assign interface {
    Target() MIRVar
}

type ValueAssign interface {
    Assign
    Source() MIROperand
}

type UnaryAssign interface {
    Assign
    Op() MIROperator
    X() MIROperand
}

type BinaryAssign interface {
    UnaryAssign
    Y() MIROperand
}

type ConditionalAssign interface {
    ValueAssign
    Cond() MIROperand
}

type CastAssign interface {
    ValueAssign
    Type() MIRType
}

type IndirectAssign interface {
    ValueAssign
}
