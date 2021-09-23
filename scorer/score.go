package scorer

import "github.com/bluekeyes/go-gitdiff/gitdiff"

type LineScoreAlgorithm interface {
	ScoreLine(line string, extension string, op gitdiff.LineOp) float64
}

type PreambleScoreAlgorithm interface {
	ScorePreamble(preamble string) float64
}

type LineScorer struct {
	LineScorer     LineScoreAlgorithm
	PreambleScorer PreambleScoreAlgorithm
}

func (s *LineScorer) SetLineScoringAlgorithm(algorithm LineScoreAlgorithm) {
	s.LineScorer = algorithm
}
func (s *LineScorer) SetPreambleScoringAlgorithm(algorithm PreambleScoreAlgorithm) {
	s.PreambleScorer = algorithm
}
