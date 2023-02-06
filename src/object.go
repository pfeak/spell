package src

import (
	"regexp"
	"strings"
)

type LCSObject struct {
	re        *regexp.Regexp // regexp wildcard
	lcsSeq    []string       // template sequence
	positions []int          // wildcard position
	splitRule string         // split rule
	label     string         // placeholder
}

func NewLCSObject(sequence []string, splitRule string, label string) *LCSObject {
	if splitRule == "" {
		splitRule = "[\\s:=,]+"
	}
	if label == "" {
		label = "<*>"
	}
	re, err := regexp.Compile(splitRule)
	if err != nil {
		panic(err)
	}

	return &LCSObject{
		re:        re,
		lcsSeq:    sequence,
		positions: make([]int, 0),
		splitRule: splitRule,
		label:     label,
	}
}

func (LCSObject *LCSObject) isPosition(idx int) bool {
	for i := range LCSObject.positions {
		if idx == i {
			return true
		}
	}
	return false
}

func (LCSObject *LCSObject) getPosition() (positions []int) {
	for i := 0; i < len(LCSObject.lcsSeq); i++ {
		if LCSObject.lcsSeq[i] == LCSObject.label {
			positions = append(positions, i)
		}
	}
	return positions
}

func (LCSObject *LCSObject) Length() int {
	return len(LCSObject.lcsSeq)
}

func (LCSObject *LCSObject) GetLCSLength(sequence []string) int {
	var (
		count, lastMatch, lenSeq int
	)
	lastMatch = -1
	lenSeq = len(sequence)
	for i := 0; i < len(LCSObject.lcsSeq); i++ {
		if LCSObject.isPosition(i) {
			continue
		}
		for j := lastMatch + 1; j < lenSeq; j++ {
			if LCSObject.lcsSeq[i] == sequence[j] {
				lastMatch = j
				count += 1
				break
			}
		}
	}
	return count
}

func (LCSObject *LCSObject) Insert(sequence []string) {
	var (
		placeholder       bool
		lastMatch, lenSeq int
		template          string
	)
	lastMatch = -1
	lenSeq = len(sequence)
	for i := 0; i < len(LCSObject.lcsSeq); i++ {
		if LCSObject.isPosition(i) {
			if !placeholder {
				template = template + LCSObject.label + " "
			}
			placeholder = true
			continue
		}
		for j := lastMatch + 1; j < lenSeq; j++ {
			if LCSObject.lcsSeq[i] == sequence[j] {
				template = template + LCSObject.lcsSeq[i] + " "
				placeholder = false
				lastMatch = j
				break
			}
			if !placeholder {
				template = template + LCSObject.label + " "
				placeholder = true
			}
		}
	}

	template = strings.TrimSpace(template)
	LCSObject.lcsSeq = LCSObject.re.Split(strings.TrimSpace(template), -1)
	LCSObject.positions = LCSObject.getPosition()
}
