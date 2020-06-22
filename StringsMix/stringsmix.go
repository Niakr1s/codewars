package stringsmix

import (
	"fmt"
	"sort"
	"strings"
)

func Mix(s1, s2 string) string {
	f1, f2 := getFreq(s1), getFreq(s2)
	diffs := SortedDiff{}
	for r := 'a'; r <= 'z'; r++ {
		cnt1, ok1 := f1[r]
		cnt2, ok2 := f2[r]
		var bestString StringID
		bestCnt := 0
		if !ok1 && !ok2 {
			continue
		}
		if ok1 && ok2 {
			if cnt1 > cnt2 {
				bestString = FirstString
				bestCnt = cnt1
			} else if cnt1 < cnt2 {
				bestString = SecondString
				bestCnt = cnt2
			} else {
				bestString = BothStrings
				bestCnt = cnt1
			}
		}
		if ok1 && !ok2 {
			bestString = FirstString
			bestCnt = cnt1
		}
		if !ok1 && ok2 {
			bestString = SecondString
			bestCnt = cnt2
		}
		if bestCnt > 1 {
			diffs = append(diffs, Diff{bestString, r, bestCnt})
		}
	}
	sort.Sort(diffs)
	str := make([]string, 0, len(diffs))
	for _, diff := range diffs {
		str = append(str, diff.String())
	}
	return strings.Join(str, "/")
}

func getFreq(s string) map[rune]int {
	res := make(map[rune]int)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res[r]++
		}
	}
	return res
}

type StringID int

const (
	FirstString StringID = iota
	SecondString
	BothStrings
)

type Diff struct {
	ID    StringID
	R     rune
	Count int
}

func (d Diff) String() string {
	id := ""
	switch d.ID {
	case FirstString:
		id = "1"
	case SecondString:
		id = "2"
	case BothStrings:
		id = "="
	}
	runes := strings.Repeat(string(d.R), d.Count)
	return fmt.Sprintf("%s:%s", id, runes)
}

type SortedDiff []Diff

func (d SortedDiff) Len() int {
	return len(d)
}

func (d SortedDiff) Less(i, j int) bool {
	if d[i].Count != d[j].Count {
		return d[i].Count > d[j].Count
	}
	if d[i].ID == d[j].ID {
		return d[i].R < d[j].R
	}
	return d[i].ID < d[j].ID
}

func (d SortedDiff) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
