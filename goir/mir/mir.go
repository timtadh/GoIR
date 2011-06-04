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

type UnaryAssign struct {
    kind kinds.Kind
    t    MIRVar
    op   kinds.UnaryOp
    x    MIRVar
}

type BinaryAssign struct {
    kind kinds.Kind
    t    MIRVar
    op   kinds.BinaryOp
    x    MIRVar
    y    MIRVar
}

// type
