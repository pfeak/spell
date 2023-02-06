package src

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type LCSMap struct {
	lcsObjects []*LCSObject   // list of LCSObject
	re         *regexp.Regexp // regexp wildcard
	splitRule  string         // split rule
	label      string         // placeholder
}

func NewLCSMap(splitRule string, label string) *LCSMap {
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

	return &LCSMap{
		lcsObjects: make([]*LCSObject, 0),
		re:         re,
		splitRule:  splitRule,
		label:      label,
	}
}

func (LCSMap *LCSMap) match(sequence []string, similarity float32) *LCSObject {
	var (
		bestMatch *LCSObject

		bestMatchLen, seqLen, objectLen, lcsLen int
	)
	seqLen = len(sequence)
	for _, object := range LCSMap.lcsObjects {
		objectLen = object.Length()
		// If similarity = 0.5, the length is less than 50%, or more than 2 times
		// If similarity = 0.1, the length is less than 10%, or more than 10 times
		if float32(objectLen) < float32(seqLen)*similarity {
			continue
		}
		if float32(objectLen) > float32(seqLen)/similarity {
			continue
		}

		lcsLen = object.GetLCSLength(sequence)

		// The value of similarity is 1~seq_len, corresponding to 100%~1% similarity
		if float32(lcsLen) > float32(seqLen)*similarity && lcsLen > bestMatchLen {
			bestMatch = object
			bestMatchLen = lcsLen
		}
	}

	return bestMatch
}

func (LCSMap *LCSMap) Train(entry string, similarity float32) (*LCSObject, error) {
	var (
		lcsObject        *LCSObject
		tmpSeq, sequence []string
	)

	if similarity < 0.01 || similarity > 1 {
		return nil, errors.New("similarity expected value [0.01,1]")
	}
	if entry == "" {
		return lcsObject, nil
	}

	tmpSeq = LCSMap.re.Split(strings.TrimSpace(entry), -1)

	for i := range tmpSeq {
		if tmpSeq[i] != "" {
			sequence = append(sequence, tmpSeq[i])
		}
	}
	lcsObject = LCSMap.match(sequence, similarity)
	if lcsObject == nil {
		lcsObject = NewLCSObject(sequence, LCSMap.splitRule, LCSMap.label)
		LCSMap.lcsObjects = append(LCSMap.lcsObjects, lcsObject)
	}

	lcsObject.Insert(sequence)
	return lcsObject, nil
}

func (LCSMap *LCSMap) Print() {
	for _, object := range LCSMap.lcsObjects {
		fmt.Printf("Template: %v\nPosition: %v\n\n", object.lcsSeq, object.getPosition())
	}
}
