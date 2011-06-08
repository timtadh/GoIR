GoIR
====

Produces a MIR (Midlevel Intermediate Representation) for Go based on the AST produced by GoAST.
The MIR is intended to be useful starting point for program analysis/optimization.

Project Status: Nothing to see here.

Todo
====

- Define semantics of the MIR [Partial]
    - Define instruction format [Done, I chose quads]
    - Define variables, types, labels etc
- Make symbol table supporting types, variables, and labels
- Start work on generator
