package sample_go_json2excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
)

func Excelize() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// ワークシートを作成する
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// セルの値を設定
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// ワークブックのデフォルトワークシートを設定します
	f.SetActiveSheet(index)
	// 指定されたパスに従ってファイルを保存します
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ExcelizeStream() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	styleID, err := f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "777777"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := sw.SetRow("A1",
		[]interface{}{
			excelize.Cell{StyleID: styleID, Value: "Data"},
			[]excelize.RichTextRun{
				{Text: "Rich ", Font: &excelize.Font{Color: "2354e8"}},
				{Text: "Text", Font: &excelize.Font{Color: "e83723"}},
			},
		},
		excelize.RowOpts{Height: 45, Hidden: false}); err != nil {
		fmt.Println(err)
		return
	}
	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, err := excelize.CoordinatesToCellName(1, rowID)
		if err != nil {
			fmt.Println(err)
			break
		}
		if err := sw.SetRow(cell, row); err != nil {
			fmt.Println(err)
			break
		}
	}
	if err := sw.Flush(); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
