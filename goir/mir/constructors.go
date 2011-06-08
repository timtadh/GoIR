package mir

import kinds

func NewLabel(label string) Label {
    return Label(&inst{kind:kinds.Label, x:MIROperand(label)})
}

type Label interface {
    MIRInst
    Label() string
}

type Assign interface {
    MIRInst
    Target() MIRVar
}

type ValueAssign interface {
    Assign
    X() MIROperand
}

type UnaryAssign interface {
    ValueAssign
    Op() MIROperator
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

type Branch interface {
    MIRInst
    Label() MIRLabel
}

type Goto interface {
    Branch
}

type Return interface {
    Branch
}

type Call interface {
    Branch
}

type ValueIf interface {
    Branch
    X() MIROperand
}

type UnaryIf interface {
    ValueIf
    Op() MIROperator
}

type BinaryIf interface {
    UnaryIf
    Y() MIROperand
}

type Parameter interface {
    MIRInst
    Param() MIRVar
    Pos() uint
    Type() MIRType
}

type Recieve interface {
    MIRInst
    Target() MIRVar
    Pos() uint
    Type() MIRType
}

