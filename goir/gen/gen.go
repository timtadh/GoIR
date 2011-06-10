package gen

import "os"
import "fmt"
import "goir/table"
import "goast/tree"

type generator struct {
    table *table.SymbolTable
}

type production func(*tree.Node) interface{}

func Generate(AST *tree.Node) map[string]interface{} {
    fmt.Println("gotcha!")
    productions := newGenerator().productions()
    productions["Package"](AST)
    return nil
}

func newGenerator() *generator {
    return &generator{
        table: table.NewSymbolTable(),
    }
}

func (self *generator) productions() map[string]production {
    var productions map[string]production

    call := func(node *tree.Node) interface{} {
        if f, ok := productions[node.Label]; ok {
            return f(node)
        }
        panic(os.NewError(fmt.Sprintf("Node '%s' does not have a handler.", node.Label)))
    }

    trycall := func(node *tree.Node, except func(node *tree.Node)) interface{} {
        defer func() {
            if err := recover(); err != nil {
                fmt.Println(err)
                except(node)
            }
        }()
        return call(node)
    }

    productions = map[string]production{
        "Package": func(node *tree.Node) interface{} {
            fmt.Println("package")
            for _, child := range node.Children {
                call(child)
            }
            return nil
        },

        "File": func(node *tree.Node) interface{} {
            fmt.Println("file")
            pack_name := call(node.Children[0])
            fmt.Println(pack_name)
            decls := call(node.Children[1])
            fmt.Println(decls)
            for _, c := range node.Children[2:] {
                fmt.Println("unprocessed ->", c.Label)
            }
            return nil
        },

        "Ident": func(node *tree.Node) interface{} {
            fmt.Println("ident")
            v := &Ident{name: node.Children[0].Label}
            if len(node.Children) == 2 {
                v.public = true
            }
            return v
        },


        "Decls": func(node *tree.Node) interface{} {
            fmt.Println("decls")
            for _, c := range node.Children {
                trycall(c,
                        func(c *tree.Node) {
                            fmt.Println("unprocessed ->", c.Label)
                        },
                )
            }
            return nil
        },
    }

    return productions
}
