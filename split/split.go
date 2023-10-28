package split

import (
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

const SEMI = 7

func SplitWithScanner(sql string) []string {
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
		if t.GetTokenType() != SEMI {
			buf.WriteString(t.GetText())
			continue
		}

		if buf.Len() == 0 {
			continue
		}

		tmp := strings.TrimSpace(buf.String())
		buf.Reset()
		if len(tmp) == 0 {
			continue
		}
		ret = append(ret, tmp)
	}
	return ret
}
