package ast

func Pipeline() {

}

type Token interface {
	Get()
}

// ValueToken is leaf node
type ValueToken [][]string

// OperatorToken is non leaf node
type OPERATOR_TYPE string

const (
	OperatorTypeAdd = "AND"
	OperatorTypeOr  = "OR"
)

func parseToAST() {

}

func calcFromAST() {

}
