package preambles

type UnimplementedPreambleScorerExample struct{}

func (s *UnimplementedPreambleScorerExample) ScorePreamble(preamble string) float64 {
	if preamble != "" {
		return 1.0
	} else {
		return 0.0
	}
}

//
////fmt.Println("scoring preamble:\n", preamble)
//left := "Subject: [PATCH]"
//right := "---"
//rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
//extracted := rx.FindAllStringSubmatch(preamble, -1)[0][0]
//
//fmt.Println("OPERATING:", extracted)
//isConventional := strings.Split(extracted, "\n")
//fmt.Println(isConventional)
//fmt.Println("\n\n")
//
