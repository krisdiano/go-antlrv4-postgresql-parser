package split

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/database-research-lab/go-antlrv4-postgresql-parser/parser"
)

type TreeShapeListener struct {
	*parser.BasePostgreSQLParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

var ErrUnexpectedEOF = errors.New("expected terminal semicolon")

func IsEndWithoutTerminal(err error) bool {
	return errors.Is(err, ErrUnexpectedEOF)
}

func SplitWithScanner(sql string) (_ []string, err error) {
	input := antlr.NewInputStream(sql)
	lexer := parser.NewPostgreSQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	stream.Fill()
	tokens := stream.GetAllTokens()

	var (
		buf strings.Builder
		ret []string
	)
	for _, t := range tokens {
		ttype := t.GetTokenType()
		if ttype == antlr.TokenEOF {
			err = ErrUnexpectedEOF
		} else if t.GetTokenType() != parser.PostgreSQLLexerSEMI {
			buf.WriteString(t.GetText())
			continue
		}

		if buf.Len() == 0 {
			err = nil
			continue
		}

		tmp := strings.TrimSpace(buf.String())
		buf.Reset()
		if len(tmp) == 0 {
			err = nil
			continue
		}
		ret = append(ret, tmp)
	}
	return ret, err
}
