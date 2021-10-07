package lines

import (
	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

type BasicLineScorer struct{}

type LineContents struct {
	Line      string
	Extension string
	Op        gitdiff.LineOp
}

func (s *BasicLineScorer) ScoreLine(line, prevLine LineContents) float64 {
	devaluedExtensions := []string{"md", "json"}
	devalueWeight := 0.4
	weight := 0.0

	if line.Line != "" && len(line.Line) > 3 {
		weight += 1.0
	}

	if contains(devaluedExtensions, line.Extension) {
		return weight * devalueWeight
	}
	return weight
}
