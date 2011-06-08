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

func NewGoto(label MIRLabel) Goto {
    return Goto(&inst{
        kind: kinds.Goto,
        z:    MIROperand(label),
    })
}

func NewReturn(label MIRLabel) Return {
    return Return(&inst{
        kind: kinds.Return,
        z:    MIROperand(label),
    })
}

func NewCall(label MIRLabel) Call {
    return Call(&inst{
        kind: kinds.Call,
        z:    MIROperand(label),
    })
}

func NewValueIf(label MIRLabel, x MIROperand) ValueIf {
    return ValueIf(&inst{
        kind: kinds.ValueIf,
        x:    x,
        z:    MIROperand(label),
    })
}

func NewUnaryIf(label MIRLabel, op MIROperator, x MIROperand) UnaryIf {
    return UnaryIf(&inst{
        kind: kinds.UnaryIf,
        op:   op,
        x:    x,
        z:    MIROperand(label),
    })
}

func NewBinaryIf(label MIRLabel, op MIROperator, x, y MIROperand) BinaryIf {
    return BinaryIf(&inst{
        kind: kinds.BinaryIf,
        op:   op,
        x:    x,
        y:    y,
        z:    MIROperand(label),
    })
}

func newParameter(kind kinds.Kind, param MIRVar, pos uint, Type MIRType) Parameter {
    return Parameter(&inst{
        kind: kind,
        x:    MIROperand(pos),
        y:    MIROperand(Type),
        z:    MIROperand(param),
    })
}

func NewInParam(param MIRVar, pos uint, Type MIRType) Parameter {
    return newParamter(kinds.InParam, param, pos, Type)
}

func NewOutParam(param MIRVar, pos uint, Type MIRType) Parameter {
    return newParamter(kinds.OutParam, param, pos, Type)
}

func newRecieve(kind kinds.Kind, target MIRVar, pos uint, Type MIRType) Recieve {
    return Recieve(&inst{
        kind: kind,
        x:    MIROperand(pos),
        y:    MIROperand(Type),
        z:    MIROperand(target),
    })
}

func NewInRecieve(target MIRVar, pos uint, Type MIRType) Recieve {
    return newRecieve(kinds.InRecieve, param, pos, Type)
}

func NewOutRecieve(target MIRVar, pos uint, Type MIRType) Recieve {
    return newRecieve(kinds.OutRecieve, param, pos, Type)
}
