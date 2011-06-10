package gen

import "os"
import "fmt"
import "strconv"
import "strings"
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
        } else {
            return productions["default"](node)
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
            return v
        },


        "Decls": func(node *tree.Node) interface{} {
            fmt.Println("decls")
            for _, c := range node.Children {
//                 trycall(c,
//                         func(c *tree.Node) {
//                             fmt.Println("unprocessed ->", c.Label)
//                         },
//                 )
                call(c)
            }
            return nil
        },

        "GenDecl": func(node *tree.Node) interface{} {
            fmt.Println("gen_decl")
            for _, c := range node.Children {
                trycall(c,
                        func(c *tree.Node) {
                            fmt.Println("spec not yet supported ", c.Label)
                        },
                )
            }
            return nil
        },


        "FuncDecl": func(node *tree.Node) interface{} {
            fmt.Println("func_decl")
            var name *Ident
            for _, c := range node.Children {
                switch c.Label {
                case "Name":
                    name = call(c.Children[0]).(*Ident)
                case "Recieve":
                case "Type":
                case "Body":
                    call(c)
                default:
                    panic(os.NewError(fmt.Sprint("Unexpected Node ", c)))
                }
            }
            if name != nil {
                fmt.Println("name ", name)
            }
            return nil
        },

        "ImportSpec": func(node *tree.Node) interface{} {
            fmt.Println("import_spec")
            var path string
            var name string
            for _, c := range node.Children {
                switch c.Label {
                case "BasicLit":
                    path = call(c).(string)
                case "Ident":
                    name = call(c).(string)
                default:
                    panic(os.NewError(fmt.Sprintf("Node '%s' does not have a handler.", node.Label)))
                }
            }
            if name == "" {
                split := strings.Split(path, "/", 0)
                if len(split) > 0 {
                    name = split[len(split)-1]
                } else {
                    name = path
                }
            }
            fmt.Println("name :", name)
            fmt.Println("path :", path)
            return nil
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
    }

    return productions
}
