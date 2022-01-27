package preambles

/**
 * Example boilerplate for a new preamble scorer algorithm
 */
type UnimplementedPreambleScorerExample struct{}

func (s *UnimplementedPreambleScorerExample) ScorePreamble(preamble string) float64 {
	return 0.0
}

/*
 * Preamble scorer that rewards commits following the conventional commits standard
 */
type ConventionCommitPreambleScorer struct{}

// Todo! figure out how to extract the conventional commit data from this preamble
func (s *ConventionCommitPreambleScorer) ScorePreamble(preamble string) float64 {
	if preamble != "" {
		return 1.0
	} else {
		return 0.0
	}
}
