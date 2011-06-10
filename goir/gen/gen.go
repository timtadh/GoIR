package gen

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
    productions = map[string]production{
        "Package": func(node *tree.Node) interface{} {
            fmt.Println("package")
            for _, child := range node.Children {
                if f, has := productions[child.Label]; has {
                    f(child)
                } else {
                    panic(child)
                }
            }
            return nil
        },

        "File": func(node *tree.Node) interface{} {
            fmt.Println("file")
            name_node := node.Children[0]
            pack_name := productions[name_node.Label](name_node)
            fmt.Println(pack_name)
            return nil
        },

        "Ident": func(node *tree.Node) interface{} {
            fmt.Println("ident")
            v := &Var{name: node.Children[0].Label}
            if len(node.Children) == 2 {
                v.public = true
            }
            return v
        },
    }

    return productions
}
