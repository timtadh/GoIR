package mir

import "goir/mir/kinds"

type inst struct {
    kind kinds.Kind
    o    MIROperator
    x    MIROperand
    y    MIROperand
    z    MIROperand
}

type MIROperand interface{}
type MIROperator interface{}
type MIRType interface{}
type MIRLabel interface{}

type MIRVar interface {
    Name() string
    String() string
    Public() bool
}

type MIRInst interface {
    String() string
    Kind() kinds.Kind
}

type Label interface {
    MIRInst
    Label() MIRLabel
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
    UnaryAssign
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

type Receive interface {
    MIRInst
    Target() MIRVar
    Pos() uint
    Type() MIRType
}
