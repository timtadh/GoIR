package kinds

type Kind uint16

const (
    Label Kind = iota // instruction label

    // Assignments (largely the same with the exception of ParAssign)
    ValueAssign       // eg. a = b
    UnaryAssign       // eg. a = - b
    BinaryAssign      // eg. a = b + c
    ConditionalAssign // eg. a = b if c
    CastAssign        // eg. a = type(b)
    IndirectAssign    // eg. *a = b

    // Branch Statements
    Goto     // eg. goto label
    BinaryIf // eg. if a op b goto label
    UnaryIf  // eg. if op a goto label
    ValueIf  // eg. if op goto label
    Call     // eg. call name
    Return   // eg. return

    // Parameter passing
    // Muchnick's Call contains the arguments it seems simpler to break them out
    InParam    // eg. inparam a 1 type               | f := func(a, b int) (c, d int)
    OutParam   // eg. outparam a 1 type              | inparam a, 1, val ; inparam b, 2, val
    ReceiveIn  // eg. receive a 1 type               | call f <recievein ... outparam ... 2>
    ReceiveOut // eg. receive a 1 type               | recieveout c, 1, val ; receiveout d, 2
)
