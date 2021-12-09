package main

import (
	"strconv"
	"strings"
)

func (l *Line) Solve() int {
	actualDefs := make(map[int]Set)
	for _, s := range l.Signals {
		// digits 1, 4, 7 and 8 have a unique number of signals so we can deduce those straight away
		switch len(s) {
		case 2:
			actualDefs[1] = NewSet(s)
		case 3:
			actualDefs[7] = NewSet(s)
		case 4:
			actualDefs[4] = NewSet(s)
		case 7:
			actualDefs[8] = NewSet(s)
		}
	}

	correctToIncorrect := make(map[string]string)
	partialCorrections := make(map[string]string)
	correctToIncorrect["a"] = actualDefs[7].Subtract(actualDefs[1]).Character()
	partialCorrections["bd"] = actualDefs[4].Subtract(actualDefs[1]).String()

	// we can find 'g' by subtracting 1, 4 and 7 from each signal. If the signal only has 1 rune remaining it's 'g'
	for _, s := range l.Signals {
		g := NewSet(s).Subtract(actualDefs[1]).Subtract(actualDefs[4]).Subtract(actualDefs[7])
		if len(g) == 1 {
			correctToIncorrect["g"] = g.Character()
			break
		}
	}

	for _, s := range l.Signals {
		if len(s) != 6 {
			// filter out all except 6,9
			continue
		}
		e := NewSet(s).Subtract(actualDefs[1]).Subtract(actualDefs[4]).Subtract(actualDefs[7]).Subtract(NewSet(correctToIncorrect["g"]))
		if len(e) == 1 {
			minus1 := NewSet(s).Subtract(actualDefs[1])
			if len(minus1) == len(NewSet(s))-2 {
				// this is 9, only thing left is 6
				continue
			}
			correctToIncorrect["e"] = e.Character()
			break
		}
	}

	for _, s := range l.Signals {
		d := NewSet(s).Subtract(actualDefs[7]).Subtract(NewSet(correctToIncorrect["g"], correctToIncorrect["e"]))
		if len(d) == 1 {
			// we have either 0, 2 or 3
			// if minusing 1 reduces s by 1, we have '2', and can map d and infer b
			sMinus1 := NewSet(s).Subtract(actualDefs[1])
			if len(sMinus1) == len(s)-1 {
				correctToIncorrect["d"] = d.Character()
				if !strings.Contains(partialCorrections["bd"], d.Character()) {
					panic("'db' should be computable")
				}
				correctToIncorrect["b"] = strings.Trim(partialCorrections["bd"], d.String())
				break
			}
		}
	}

	// distinguish 'c' and 'f'... find 5 to deduce 'f'
	for _, s := range l.Signals {
		if len(s) != 5 {
			// this filters out all numbers except 2, 3 and 5
			continue
		}

		f := NewSet(s).Subtract(NewSet(correctToIncorrect["a"], correctToIncorrect["b"], correctToIncorrect["d"], correctToIncorrect["g"]))
		if len(f) == 1 {
			correctToIncorrect["f"] = f.Character()
			c := actualDefs[1].Subtract(f)
			if len(c) != 1 {
				panic("this should have solved 'c'")
			}
			correctToIncorrect["c"] = c.Character()
		}
	}

	incorrectToCorrect := make(map[string]string, len(correctToIncorrect))
	for k, v := range correctToIncorrect {
		incorrectToCorrect[v] = k
	}

	var sb strings.Builder
	for _, d := range l.Digits {
		mapped := NewSet()
		for i := range d {
			mapped[string(incorrectToCorrect[string(d[i])][0])] = struct{}{}
		}
		for actual, set := range digitDefinitions {
			if set.Equals(mapped) {
				sb.WriteString(actual)
				break
			}
		}
	}

	answer, _ := strconv.Atoi(sb.String())

	return answer
}

func Part2(input []Line) int {
	answer := 0
	for _, line := range input {
		answer += line.Solve()
	}

	return answer
}
