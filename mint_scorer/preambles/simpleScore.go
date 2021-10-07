package preambles

type UnimplementedPreambleScorerExample struct{}

func (s *UnimplementedPreambleScorerExample) ScorePreamble(preamble string) float64 {
	return 0.0
}
