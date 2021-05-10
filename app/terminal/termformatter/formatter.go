package termformatter

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
)

type BlockFormatter struct {
	blocks []BlockDef
	io     app.Output
	width  int
}

type BlockDef struct {
	lastLineFormatter FormatLineCB
	lineFormatter     FormatLineCB
	width             int
}

func NewBlockDef(width int, lineFormatter FormatLineCB, lastLineFormatter FormatLineCB) BlockDef {
	if width == 0 {
		panic("width is expected")
	}
	if lineFormatter == nil {
		panic("lineFormatter is expected")
	}
	if lastLineFormatter == nil {
		lastLineFormatter = lineFormatter
	}
	return BlockDef{
		lastLineFormatter: lastLineFormatter,
		lineFormatter:     lineFormatter,
		width:             width,
	}
}

func NewBlockFormatter(io app.Output, width int, blocks ...BlockDef) BlockFormatter {
	if width > maxLineLength {
		panic(fmt.Errorf("maximum width is %d and you try create blocks with width %d", maxLineLength, width))
	}
	sum := 0
	for _, block := range blocks {
		sum += block.width
	}
	if sum != width {
		panic(fmt.Errorf("expected blocks width equals to %d", width))
	}
	return BlockFormatter{
		io:     io,
		blocks: blocks,
		width:  width,
	}
}

func (formatter BlockFormatter) PrintBlocks(contents ...string) {
	if len(contents) != len(formatter.blocks) {
		panic(fmt.Errorf("content count and block count must be the same"))
	}
	contentBlocks := make([][][]string, len(contents))
	// prepare lines
	for i, content := range contents {
		width := formatter.blocks[i].width
		contentWords := SeparateWords(content)
		lines := [][]string{}
		for len(contentWords) > 0 {
			var line []string
			line, contentWords = SeparateLines(contentWords, width)
			lines = append(lines, line)
		}
		contentBlocks[i] = lines
	}
	lineNumber := 0
	for hasMore := true; hasMore; lineNumber++ {
		hasMore = false
		for i, block := range contentBlocks {
			var lineFormatter FormatLineCB
			def := formatter.blocks[i]
			if len(block) <= lineNumber {
				formatter.io.Printf(emptyLine[:def.width])
				continue
			}
			if len(block)-1 > lineNumber {
				hasMore = true
			}
			line := block[lineNumber]
			if len(block)-1 == lineNumber {
				lineFormatter = def.lastLineFormatter
			} else {
				lineFormatter = def.lineFormatter
			}
			content := lineFormatter(line, def.width)
			formatter.io.Printf(content)
		}
		formatter.io.Printf("\n")
	}
}

func (formatter BlockFormatter) Println() {
	formatter.io.Printf("\n")
}
