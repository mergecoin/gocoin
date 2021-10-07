package diffstream

import (
	"fmt"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

type Statistic struct {
	Edits     int
	Deletes   int
	Additions int
	Extension string
	Name      string
	Fragments []*gitdiff.TextFragment
}

func (s *Statistic) applyName(name string) {
	s.Name = name
}

func (s *Statistic) applyExtension(extension string) {
	s.Extension = extension
}

func (s *Statistic) applyFragments(fragments []*gitdiff.TextFragment) {
	//fmt.Println("prev", len(s.Fragments))
	s.Fragments = append(s.Fragments, fragments...)
	//fmt.Println("after", len(s.Fragments))
}

func (s *Statistic) incrementEdit() {
	s.Edits++
}

func (s *Statistic) incrementAddition() {
	s.Additions++
}

func (s *Statistic) incrementDelete() {
	s.Deletes++
}

type DiffInfo struct {
	Data     map[string]*Statistic
	Preamble string
}

func (s *DiffInfo) AddDiffPreamble(preamble string) {
	s.Preamble = preamble
}

func (s *DiffInfo) InitFileStatistic(file string) {
	_, ok := s.Data[fmt.Sprintf(file)]
	if !ok {
		s.Data[fmt.Sprintf(file)] = &Statistic{}
	}
}

func (s *DiffInfo) IncrementEdits(file string) {
	if v, found := s.Data[file]; found {
		v.incrementEdit()
	}
}

func (s *DiffInfo) IncrementAdditions(file string) {
	if v, found := s.Data[file]; found {
		v.incrementAddition()
	}
}

func (s *DiffInfo) IncrementDelete(file string) {
	if v, found := s.Data[file]; found {
		v.incrementDelete()
	}
}

func (s *DiffInfo) ApplyName(file string) {
	if v, found := s.Data[file]; found {
		v.applyName(file)
	}
}

func (s *DiffInfo) AddFragments(file string, fragments []*gitdiff.TextFragment) {
	if v, found := s.Data[file]; found {
		v.applyFragments(fragments)
	}
}

func (s *DiffInfo) AddExtension(file string, extension string) {
	if v, found := s.Data[file]; found {
		v.applyExtension(extension)
	}
}
