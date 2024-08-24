package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {

	type pattern struct {
		name, pattern string
	}
	allPatterns := []pattern{
		// {name: `all enterprise partitions`, pattern: `*`}, // refer `[^t]*`
		{name: `[internal]all enterprise partitions`, pattern: `[^t]*`},
		// {name: `all enterprise partitions recursive`, pattern: `**`}, // refer `[^t]**`
		{name: `[internal]all enterprise partitions recursive`, pattern: `[^t]**`},
		{name: `all enterprise partition's nested partitions except default'`, pattern: `*/*`},
		{name: `all enterprise recursive partition's nested partitions except default'`, pattern: `*/**`},
		{name: `[internal]all enterprise partition's nested partitions`, pattern: `[^t]*/*`},
		{name: `all tenant's default partitions'`, pattern: `t*`},
		{name: `all tenant's nested partitions except their default'`, pattern: `t*/*`},
		{name: `all tenant1's nested partitions except his default`, pattern: `t1/*`},
		{name: `all tenant1's partitions including default'`, pattern: `t1*`},
	}

	allPartitions := []string{
		"",
		"ep1",
		"ep1/esp1",
		"ep1/esp2",
		"ep1/esp1/essp1",
		"ep2",
		"t1",
		"t1/p1",
		"t1/p1/sp1",
		"t1/p1/sp1/ssp1",
		"t1/p2",
		"t2",
		"t2/p1",
		"t2/p2",
	}

	tw := table.NewWriter()
	tw.SetCaption(color.MagentaString("%T", "Partition Filter"))
	tw.AppendHeader([]any{"Description", "Pattern", "Matched", "Not Matched"})
	tw.Style().Title.Align = text.AlignCenter
	tw.SetColumnConfigs([]table.ColumnConfig{

		// {
		//	Name:     "Field Name",
		//	Colors:   text.Colors{text.FgBlue},
		//	WidthMin: 10,
		//	Align:    text.AlignCenter,
		// },
	})
	tw.Style().Color.Header = text.Colors{text.FgGreen}
	// tw.SetAutoIndex(true)
	for _, pat := range allPatterns {
		var matchedL []string
		var unmatchedL []string
		for _, k := range allPartitions {
			// matched, err := glob.Match(k, pat.pattern)
			matched := GlobBytes(globParts(pat.pattern), []byte(k))

			if !matched {
				unmatchedL = append(unmatchedL, k)
				// fmt.Println(k, color.RedString("not matched"), p)
			} else {
				matchedL = append(matchedL, k)
				// fmt.Println(k, color.GreenString("matched"), p)
			}
		}
		// fmt.Println(name, p, "matched", matchedL, "unmatched", unmatchedL)
		tw.AppendRow([]any{pat.name, pat.pattern, matchedL, unmatchedL})
	}

	fmt.Println(tw.Render())
}

var glob = "*"

func index(s, pattern string) int {
	if len(pattern) > len(s) {
		return -1
	}

	patternLen := len(pattern)
	for i := 0; i < len(s)-patternLen+1; i++ {
		ss := s[i : i+patternLen]
		if strings.EqualFold(ss, pattern) {
			return i
		}
	}
	return -1
}

func hasSuffix(s, pattern string) bool {
	if len(pattern) > len(s) {
		return false
	}
	return index(s[len(s)-len(pattern):], pattern) == 0
}

func GlobBytes(globParts [][]byte, s []byte) bool {
	// if globParts is 0 we assume that query is *
	if len(globParts) == 0 {
		return true
	}

	remaining := s
	for _, part := range globParts {
		remaining = trimPart(part, remaining)
		if remaining == nil {
			return false
		}
	}
	return true
}

func globParts(p string) [][]byte {
	s := bytes.Split([]byte(p), []byte(glob))
	res := [][]byte{}
	for _, n := range s {
		if len(n) > 0 {
			res = append(res, n)
		}
	}
	return res
}

func extractRune(s []byte, i int) (rune, int) {
	if s[i] < utf8.RuneSelf {
		return rune(s[i]), 1
	}
	return utf8.DecodeRune(s[i:])
}

func trimPart(part, s []byte) []byte {
	if len(part) == 0 {
		return s
	}

	lower, pstart := extractRune(part, 0)
	upper := unicode.ToUpper(lower)

	for i := 0; i < len(s); {
		r, size := extractRune(s, i)
		i += size
		if r != lower && r != upper {
			continue
		}
		if pstart >= len(part) {
			return s[i:]
		}
		potentialStart := false
		si, pi := i, pstart
		for si < len(s) && pi < len(part) {
			r, size = extractRune(s, si)
			si += size
			if !potentialStart {
				i = si
				potentialStart = r == lower || r == upper
			}

			pr, psize := extractRune(part, pi)
			pi += psize
			if pr == r || unicode.ToUpper(pr) == r {
				if pi >= len(part) {
					return s[si:]
				}
				continue
			}
			if !potentialStart {
				break
			}
			potentialStart = false
			si, pi = i, pstart
		}
	}
	return nil
}

func hasSuffixBytes(s, pattern []byte) bool {
	if len(pattern) > len(s) {
		return false
	}
	return bytes.EqualFold(s[len(s)-len(pattern):], pattern)
}
