package parser

type VariableType int

const (
	Mozzarella VariableType = 0
	Gorgonzola VariableType = 1
	Parmesan   VariableType = 2
	Milk       VariableType = 3
	Ricotta    VariableType = 4
	Tuple      VariableType = 5
)

func (v VariableType) String() string {
	switch v {
	case Mozzarella:
		return "Mozzarella"
	case Gorgonzola:
		return "Gorgonzola"
	case Parmesan:
		return "Parmesan"
	case Milk:
		return "Milk"
	case Ricotta:
		return "Ricotta"
	case Tuple:
		return "Tuple"
	default:
		panic("Run Time Panic: variable type unsupported")
	}
}

type VariableContainer interface {
	GetVariableType() VariableType
}

type MozzarellaVariable struct {
	Value string
}

func (*MozzarellaVariable) GetVariableType() VariableType {
	return Mozzarella
}

type GorgonzolaVariable struct {
	Value float64
}

func (*GorgonzolaVariable) GetVariableType() VariableType {
	return Gorgonzola
}

type ParmesanVariable struct {
	Value int
}

func (*ParmesanVariable) GetVariableType() VariableType {
	return Parmesan
}

type MilkVariable struct {
	Value bool
}

func (*MilkVariable) GetVariableType() VariableType {
	return Milk
}

type RicottaVariable struct{}

func (*RicottaVariable) GetVariableType() VariableType {
	return Ricotta
}

type TupleVariableType struct {
	Variables []VariableContainer
}

func (*TupleVariableType) GetVariableType() VariableType {
	return Tuple
}

type Variable struct {
	Name  string
	Value VariableContainer
}

var NullVariableContainer VariableContainer = &RicottaVariable{}
var NullVariable = Variable{"", NullVariableContainer}
