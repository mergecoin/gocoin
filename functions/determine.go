package functions

import (
	"github.com/ventureharbour/gocoin/diffscanner/diffstream"
	"github.com/ventureharbour/gocoin/retrieve"
	"github.com/ventureharbour/gocoin/scorer/lines"
	"github.com/ventureharbour/gocoin/scorer/preambles"
)

// Determines an amount of mergecoing for a given PR
func Determine(org, project, token string, pull int) float64 {
	stream := diffstream.NewDiffStream(retrieve.Retrieve(
		org,
		project,
		pull,
		token))

	stream.InitializeData()

	return stream.GenerateScore(&lines.UnimplementedLineScorerExample{}, &preambles.UnimplementedPreambleScorerExample{})
}
