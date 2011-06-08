package mir

import "kinds"

func NewLabel(label string) Label {
    return Label(&inst{kind: kinds.Label, z: MIROperand(label)})
}


func NewValueAssign(target MIRVar, x MIROperand) ValueAssign {
    return ValueAssign(&inst{
        kind: kinds.ValueAssign,
        x:    x,
        z:    MIROperand(target),
    })
}

func NewUnaryAssign(op MIROperator, target MIRVar, x MIROperand) UnaryAssign {
    return UnaryAssign(&inst{
        kind: kinds.UnaryAssign,
        o:    op,
        x:    x,
        z:    MIROperand(target),
    })
}

func NewBinaryAssign(op MIROperator, target MIRVar, x, y MIROperand) BinaryAssign {
    return BinaryAssign(&inst{
        kind: kinds.BinaryAssign,
        o:    op,
        x:    x,
        y:    y,
        z:    MIROperand(target),
    })
}

func NewConditionalAssign(cond_op MIROperator,
target MIRVar,
x, cond MIROperand) ConditionalAssign {
    return ConditionalAssign(&inst{
        kind: kinds.ConditionalAssign,
        o:    cond_op,
        x:    x,
        y:    cond,
        z:    MIROperand(target),
    })
}

func NewCastAssign(cond_op MIROperator, target MIRVar, x MIROperand, Type MIRType) CastAssign {
    return CastAssign(&inst{
        kind: kinds.CastAssign,
        o:    cond_op,
        x:    x,
        y:    MIROperand(Type),
        z:    MIROperand(target),
    })
}

func NewIndirectAssign(target MIRVar, x MIROperand) IndirectAssign {
    return IndirectAssign(&inst{
        kind: kinds.IndirectAssign,
        x:    x,
        z:    MIROperand(target),
    })
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
