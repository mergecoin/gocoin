package preambles

type UnimplementedPreambleScorerExample struct{}

func (s *UnimplementedPreambleScorerExample) ScorePreamble(preamble string) float64 {
	return 1.0
}
