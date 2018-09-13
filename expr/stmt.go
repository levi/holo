package expr

type Stmt interface {}


type Expression struct {
    expression Expr
}

func NewExpression(expression Expr) Expression {
    return Expression{
        expression,
    }
}

type Print struct {
    expression Expr
}

func NewPrint(expression Expr) Print {
    return Print{
        expression,
    }
}

