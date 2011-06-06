package mir

import "goir/mir/kinds"

type MIRInst interface {
    Kind() kinds.Kind
}

type MIRVar interface {
    Name() string
    Size() uint32    // size in bytes
    BitSize() uint32 // size in bits
}

type MIRConst interface {
    Value() string
}

type MIRType interface {
    Name() string
}

type MIROperand interface { // eg. Var, Const, or TypeName
    Name() string
}

type Label struct {
    kind kinds.Kind
    name string
}

type Assignment interface {
    MIRInst
    Targ() MIRVar
    Op() kinds.UnaryOp
    X() MIRVar
}

type ValAssign struct {
    kind kinds.Kind
    t    MIRVar
    x    MIROperand
}

type UnaryAssign struct {
    kind kinds.Kind
    t    MIRVar
    op   kinds.UnaryOp
    x    MIROperand
}

type BinaryAssign struct {
    kind kinds.Kind
    t    MIRVar
    op   kinds.BinaryOp
    x    MIROperand
    y    MIROperand
}

type CondAssign struct {
    kind kinds.Kind
    t    MIRVar
    cond MIRVar
    x    MIROperand
}

type CastAssign struct {
    kind kinds.Kind
    t    MIRVar
    typ  MIRType
    x    MIROperand
}

type IndirAssign struct {
    kind kinds.Kind
    t    MIRVar
    x    MIROperand
}

type ParAssign struct {
    kind kinds.Kind
    t    MIRVar
    u    MIRVar
    x    MIROperand
    y    MIROperand
}


// type
