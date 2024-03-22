package util

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PPSize(obj any) {
	tw := table.NewWriter()
	tw.SetCaption(color.MagentaString("%T", obj))
	tw.AppendHeader([]any{"Field Name", "Offset", "Size", "Byte Map"})
	tw.Style().Title.Align = text.AlignCenter
	tw.SetColumnConfigs([]table.ColumnConfig{

		//{
		//	Name:     "Field Name",
		//	Colors:   text.Colors{text.FgBlue},
		//	WidthMin: 10,
		//	Align:    text.AlignCenter,
		//},
	})
	tw.Style().Color.Header = text.Colors{text.FgGreen}
	tw.SetAutoIndex(true)
	t := reflect.TypeOf(obj)
	//prevOff := 0

	expectedBytes := 0
	wastedBytes := 0
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fOffset := int(f.Offset)
		fSize := int(f.Type.Size())
		expectedBytes += fSize
		wordOffset := 0
		byteMap := ""
		for i := 0; i < fOffset%8; i++ {
			//byteMap += color.CyanString("□ ")
			byteMap += color.CyanString("  ")
			wordOffset = (wordOffset + 1) % 8
			if wordOffset == 0 {
				byteMap += "\n"
			}
		}
		j := fOffset
		for ; j < fOffset+fSize; j++ {
			byteMap += color.GreenString("□ ")
			wordOffset = (wordOffset + 1) % 8
			if wordOffset == 0 {
				byteMap += "\n"
			}
		}
		nextOffset := 0
		if i < t.NumField()-1 {
			nextOffset = int(t.Field(i + 1).Offset)
		} else {
			if j%8 != 0 {
				nextOffset = j + (8 - (j % 8))
			}
		}
		for ; j < nextOffset; j++ {
			wastedBytes++
			byteMap += color.RedString("□ ")
			wordOffset = (wordOffset + 1) % 8
			if wordOffset == 0 {
				byteMap += "\n"
			}
		}
		tw.AppendRow([]any{f.Name, f.Offset, f.Type.Size(), byteMap})
	}

	str := fmt.Sprintf("Expected: %d Bytes, ", expectedBytes)
	if wastedBytes > 8 {
		str += color.RedString("Wasted  : %d Bytes", wastedBytes)
	} else {
		str += fmt.Sprintf("Wasted  : %d Bytes", wastedBytes)
	}
	tw.AppendFooter([]any{"", "",
		str,
		fmt.Sprintf("Actual: %d Bytes", t.Size()),
	})

	fmt.Println(tw.Render())

}
