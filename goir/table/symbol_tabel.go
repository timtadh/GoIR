package table

type SymbolTable struct {
    table map[string]interface{}
    parent *SymbolTable
}

// Constructor for a new SymbolTable. You should only call this once, when constructing
// the root of the table. All other times you probably want to use Push()
func NewSymbolTable() *SymbolTable {
    return &SymbolTable{
        table: make(map[string]interface{}),
    }
}

// Get a new symbol table with the current table as the parent.
func (self *SymbolTable) Push() *SymbolTable {
    t := NewSymbolTable()
    t.parent = self
    return t
}

// Get the parent off the current table.
func (self *SymbolTable) Pop() *SymbolTable {
    return self.parent
}

// Get a key from the table, search recursively up to the root symbol table.
func (self *SymbolTable) Get(key string) (interface{}, bool) {
    if item, has := self.table[key]; has {
        return item, true
    }
    if self.parent != nil {
        return self.parent.Get(key)
    }
    return nil, false
}

// Set a key in the current table.
func (self *SymbolTable) Set(key string, item interface{}) {
    self.table[key] = item
}

// Delete a key in the *current* table. This does not delete keys in parent tables.
func (self *SymbolTable) Del(key string) {
    self.table[key] = nil, false
}

func (self *SymbolTable) String() string {
    return "need to actually implement this method"
}
