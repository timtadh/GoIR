package gen

import "os"
import "fmt"
import "strconv"
// import "strings"
import "goir/table"
import "goast/tree"
import "goir/mir"

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
        } else {
            return productions["default"](node)
        }
        panic(os.NewError(fmt.Sprintf("Node '%s' does not have a handler.", node.Label)))
    }

    tmpvar := func() mir.MirVar

//     trycall := func(node *tree.Node, except func(node *tree.Node)) interface{} {
//         defer func() {
//             if err := recover(); err != nil {
//                 fmt.Println(err)
//                 except(node)
//             }
//         }()
//         return call(node)
//     }

    productions = map[string]production{
        "default": func(node *tree.Node) interface{} {
            fmt.Println("default ", node.Label)
            ret := make([]interface{}, 0, len(node.Children))
            for _, c := range node.Children {
                ret = append(ret, call(c))
            }
            return ret
        },

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
            fmt.Println(v)
            return v
        },

        "BasicLit": func(node *tree.Node) interface{} {
            fmt.Println("basic_lit")
            type_ := node.Children[0].Label
            value := node.Children[1].Label
            switch type_ {
            case "STRING":
                return value[1:len(value)-1]
            case "INT":
                if i, err := strconv.Atoi64(value); err == nil {
                    fmt.Println(i)
                    return i
                } else {
                    panic(err)
                }
            default:
                panic(os.NewError(fmt.Sprintf("Basic Lit type '%s' is not supported.", type_)))
            }
            return nil
        },

        "BinaryExpr": func(node *tree.Node) interface{} {
            fmt.Println(node.Label)
            op := node.Children[0].Label
            x := call(node.Children[1]).(mir.MIROperand)
            y := call(node.Children[2]).(mir.MIROperand)
            inst := mir.NewBinaryAssign(op, nil, x, y)
            return inst
        },
    }

    return productions
}
