package diffstream

import (
	"bufio"
	diffstream "github.com/ventureharbour/gocoin/diffscanner/diffinfo"
	"github.com/ventureharbour/gocoin/mint_scorer"
	"log"
	"regexp"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

type DiffStream struct {
	*bufio.Scanner
	readData *strings.Reader
	Info     *diffstream.DiffInfo
}

func NewDiffStream(data []byte) *DiffStream {
	readData := strings.NewReader(string(data))

	return &DiffStream{
		bufio.NewScanner(readData),
		readData,
		nil,
	}
}

func (s *DiffStream) InitializeData() {
	files, preamble, _ := gitdiff.Parse(s.readData)

	rgx, err := regexp.Compile("([^\\.]+$)")
	if err != nil {
		log.Fatal(err)
	}

	stats := &diffstream.DiffInfo{
		Preamble: preamble,
		Data:     make(map[string]*diffstream.Statistic),
	}

	for _, file := range files {
		stats.InitFileStatistic(file.NewName)
		stats.ApplyName(file.NewName)
		if !file.IsBinary {
			stats.AddExtension(file.NewName, string(rgx.Find([]byte(file.NewName))))
			stats.AddFragments(file.NewName, file.TextFragments)
			for _, fragment := range file.TextFragments {
				for _, line := range fragment.Lines {
					if !line.New() {
						stats.IncrementEdits(file.NewName)
					} else {
						if line.Op == gitdiff.OpAdd {
							stats.IncrementAdditions(file.NewName)
						}
						if line.Op == gitdiff.OpDelete {
							stats.IncrementDelete(file.NewName)
						}
					}
				}
			}
		}
	}
	s.Info = stats
}

func (s *DiffStream) GenerateScore(lineAlgorithm mint_scorer.LineScoreAlgorithm, preambleAlgorithm mint_scorer.PreambleScoreAlgorithm) float64 {
	scoring := mint_scorer.LineScorer{}
	scoring.SetLineScoringAlgorithm(lineAlgorithm)
	scoring.SetPreambleScoringAlgorithm(preambleAlgorithm)


	total := 0.0

	for _, dataPoint := range s.Info.Data {
		for _, fragment := range dataPoint.Fragments {
			for _, line := range fragment.Lines {
				total += scoring.LineScorer.ScoreLine(line.Line, dataPoint.Extension, line.Op)
			}
		}
	}

	total += scoring.PreambleScorer.ScorePreamble(s.Info.Preamble)

	return total
}
