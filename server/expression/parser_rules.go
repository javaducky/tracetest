package expression

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Statement struct {
	Left       *Expr  `@@`
	Comparator string `@Comparator`
	Right      *Expr  `@@`
}

type Expr struct {
	Left    *Term     `@@`
	Right   []*OpTerm `@@*`
	Filters []*Filter `@@*`
}

type OpTerm struct {
	Operator string `@Operator`
	Term     *Term  `@@`
}

type Term struct {
	FunctionCall *FunctionCall `( @@`
	Array        *Array        `| @@`
	Duration     *string       `| @Duration `
	Number       *string       `| @Number `
	Attribute    *Attribute    `| @Attribute `
	Environment  *Environment  `| @Environment `
	Str          *Str          `| @(QuotedString|SingleQuotedString) )`
}

type Filter struct {
	Name string  `"|" @Ident`
	Args []*Term `@@*`
}

type FunctionCall struct {
	Name string  `@Ident`
	Args []*Term `"(" ( @@ ("," @@ )* )? ")"`
}

type Array struct {
	Items []*Term `"[" ( @@ ("," @@ )* )? "]"`
}

var languageLexer = lexer.MustStateful(lexer.Rules{
	"Root": {
		{Name: "whitespace", Pattern: `\s+`, Action: nil},
		{Name: "Punc", Pattern: `[(),|\[\]]`, Action: nil},

		{Name: "Comparator", Pattern: `!=|<=|>=|=|<|>|contains|not-contains`},
		{Name: "Operator", Pattern: `(\+|\-|\*|\/)`, Action: nil},

		{Name: "Duration", Pattern: `([0-9]+(\.[0-9]+)?)(ns|us|ms|s|m|h)`},
		{Name: "Number", Pattern: `([0-9]+(\.[0-9]+)?)`},
		{Name: "Attribute", Pattern: `attr:[a-zA-Z_0-9][a-zA-Z_0-9.]*`, Action: nil},
		{Name: "Environment", Pattern: `env:[a-zA-Z_0-9][a-zA-Z_0-9.]*`, Action: nil},
		{Name: "QuotedString", Pattern: `"(\\"|[^"])*"`, Action: nil},
		{Name: "SingleQuotedString", Pattern: `'(\\'|[^'])*'`, Action: nil},

		{Name: "Ident", Pattern: `[a-zA-Z][a-zA-Z0-9_]*`, Action: nil},
	},
})

type TermType = string

const (
	FunctionCallType TermType = "FunctionCall"
	ArrayType        TermType = "Array"
	DurationType     TermType = "Duration"
	NumberType       TermType = "Number"
	AttributeType    TermType = "Attribute"
	EnvironmentType  TermType = "Environment"
	StrType          TermType = "Str"
)

func (term *Term) Type() TermType {
	if term.Attribute != nil {
		return AttributeType
	}

	if term.Environment != nil {
		return EnvironmentType
	}

	if term.FunctionCall != nil {
		return FunctionCallType
	}

	if term.Array != nil {
		return ArrayType
	}

	if term.Duration != nil {
		return DurationType
	}

	if term.Number != nil {
		return NumberType
	}

	if term.Str != nil {
		return StrType
	}

	return ""
}

func (expr *Expr) GetTermsByType(termType TermType) []*Term {
	terms := []*Term{}

	if expr.Left.Type() == StrType {
		for _, arg := range expr.Left.Str.Args {
			terms = append(terms, arg.GetTermsByType(termType)...)
		}
	}

	if expr.Left.Type() == termType {
		terms = append(terms, expr.Left)
	}

	for _, optTerm := range expr.Right {
		if optTerm.Term.Type() == termType {
			terms = append(terms, optTerm.Term)
		}
	}

	return terms
}
