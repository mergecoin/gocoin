package lines

import "github.com/bluekeyes/go-gitdiff/gitdiff"

type UnimplementedLineScorerExample struct{}

func (s *UnimplementedLineScorerExample) ScoreLine(line string, extension string, op gitdiff.LineOp) float64 {
	return 1.0
}
