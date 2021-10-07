package mint_scorer

import (
	"github.com/ventureharbour/gocoin/mint_scorer/lines"
)

type LineScoreAlgorithm interface {
	ScoreLine(contents, prevLine lines.LineContents) float64
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
