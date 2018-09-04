package holo

type Expr interface {}

type Binary struct {
    left Expr{}
    operation Token
    right Expr{}
}

func NewBinary(left Expr{}, operation Token, right Expr{}) *Binary {
    n = new(Binary)
    n.left = left
    n.operation = operation
    n.right = right
    return n
}

type Grouping struct {
    expression Expr{}
}

func NewGrouping(expression Expr{}) *Grouping {
    n = new(Grouping)
    n.expression = expression
    return n
}

type Literal struct {
    value Object
}

func NewLiteral(value Object) *Literal {
    n = new(Literal)
    n.value = value
    return n
}

type Unary struct {
    operation Token
    right Expr{}
}

func NewUnary(operation Token, right Expr{}) *Unary {
    n = new(Unary)
    n.operation = operation
    n.right = right
    return n
}

