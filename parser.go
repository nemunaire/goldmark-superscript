package superscript

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type superscriptDelimiterProcessor struct {
}

func (p *superscriptDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '^'
}

func (p *superscriptDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *superscriptDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return &Node{}
}

var defaultEmphasisDelimiterProcessor = &superscriptDelimiterProcessor{}

type superscriptParser struct {
}

var defaultSuperscriptParser = &superscriptParser{}

func (s *superscriptParser) Trigger() []byte {
	return []byte{'^'}
}

func (s *superscriptParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 1, defaultEmphasisDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}
