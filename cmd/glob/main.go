package main

import (
	"fmt"

	"github.com/fatih/color"
	glob "github.com/ganbarodigital/go_glob"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	// data := []map[string][]string{
	// 	{"t1/*": {"t1/", "t1/p1"}},
	// 	{"t1*": {"t1", "t1/p1"}},
	// 	{"/*": {"/ep1", "t1/p1"}},
	// 	{"*/*": {"/ep1", "t1/p1"}},
	// 	{`t1/[^(p1)]*`, pattern: {"t1/", "t1/p1"}},
	// }
	// for _, d := range data {
	// 	for k, v := range d {
	// 		fmt.Println("========================")
	// 		for _, s := range v {
	// 			matched, err := filepath.Match(k, s)
	// 			if err != nil {
	// 				fmt.Println(k, "not matched", s, " : ", err)
	// 			}
	// 			if !matched {
	//
	// 				fmt.Println(k, color.RedString("not matched"), s)
	// 			} else {
	// 				fmt.Println(k, color.GreenString("matched"), s)
	// 			}
	// 		}
	// 	}
	// }
	//
	// fmt.Println("####################################")
	//
	// data2 := []map[string][]string{
	// 	{"t1/*": {"t1", "t1/p1"}},
	// 	{"t1*": {"t1", "t1/p1"}},
	// 	{"t2*": {"t2", "t2/p1"}},
	// 	{"t*": {"t1", "t2", "t1/p1", "t2/p1"}},
	// 	{"t*/*": {"t1", "t2", "t1/p1", "t2/p1"}},
	// 	{"[^t]*": {"ep1", "t1"}}, // user given * needs to be modified as [^t]* for evaluation
	// 	{"**": {"ep1", "t1/p1"}},
	// }
	// for _, d := range data2 {
	// 	for k, v := range d {
	// 		fmt.Println("========================")
	// 		for _, s := range v {
	// 			matched, err := glob.Match(s, k)
	// 			if err != nil {
	// 				fmt.Println(k, "not matched", s, " : ", err)
	// 			}
	// 			if !matched {
	//
	// 				fmt.Println(k, color.RedString("not matched"), s)
	// 			} else {
	// 				fmt.Println(k, color.GreenString("matched"), s)
	// 			}
	// 		}
	// 	}
	// }

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
		{name: `all tenant1's nested partitions except his default`, pattern: `t1/*`}, // this doesn't work as expected
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

	// for _, p := range allPartitions {
	// 	for _, k := range allPatterns {
	// 		matched, err := glob.Match(p, k)
	// 		if err != nil {
	// 			fmt.Println(k, "not matched", p, " : ", err)
	// 		}
	// 		if !matched {
	// 			fmt.Println(k, color.RedString("not matched"), p)
	// 		} else {
	// 			fmt.Println(k, color.GreenString("matched"), p)
	// 		}
	// 	}
	// }

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
			matched, err := glob.Match(k, pat.pattern)
			// p, b, err := glob.MatchShortestPrefix(k, pat.pattern)
			// fmt.Println("matched", p, b, err, k, pat.pattern, k[:p])
			// matched := err != nil && b && p == len(k)
			if err != nil {
				unmatchedL = append(unmatchedL, k)
				// fmt.Println(k, "not matched", p, " : ", err)
			}
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
