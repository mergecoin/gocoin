package diffstream

import (
	"bufio"
	"fmt"
	"github.com/ventureharbour/gocoin/config"
	diffstream "github.com/ventureharbour/gocoin/diffscanner/diffinfo"
	"github.com/ventureharbour/gocoin/mint_scorer"
	"github.com/ventureharbour/gocoin/mint_scorer/lines"
	"github.com/ventureharbour/gocoin/utils"
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

func (s *DiffStream) GenerateScore(lineAlgorithm mint_scorer.LineScoreAlgorithm, preambleAlgorithm mint_scorer.PreambleScoreAlgorithm, config config.DeterminationConfig) (float64, float64) {
	scoring := mint_scorer.LineScorer{}
	scoring.SetLineScoringAlgorithm(lineAlgorithm)
	scoring.SetPreambleScoringAlgorithm(preambleAlgorithm)

	total := 0.0
	preambleExtra := 0.0
	prevLine := lines.LineContents{}

	ignoredFileNames := config.Ignored.Names

	fmt.Println(ignoredFileNames)

	for _, dataPoint := range s.Info.Data {
		if !utils.Includes(dataPoint.Name, ignoredFileNames...) {
			for _, fragment := range dataPoint.Fragments {
				for _, line := range fragment.Lines {
					// indicates that this Line is not a context Line in the patch
					if line.Op != 0 {
						contents := lines.LineContents{
							Line:      strings.TrimSpace(line.Line),
							Extension: dataPoint.Extension,
							Op:        line.Op,
						}
						total += scoring.LineScorer.ScoreLine(contents, prevLine)
						prevLine = contents
					}
				}
			}
		}
		preambleExtra += scoring.PreambleScorer.ScorePreamble(s.Info.Preamble)
	}
	return total, preambleExtra
}
