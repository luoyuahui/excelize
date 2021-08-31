package excelize_test

import (
	"log"
	"os"
	"testing"

	"github.com/luoyuahui/excelize/v2"
)

func Test_WriteStream(t *testing.T) {
	f, err := os.OpenFile("test.xlsx", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ef := excelize.NewStream(f)
	ssw, err := ef.NewStreamWriter("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	ssw.SetColWidth(2, 3, 80)
	for i := 0; i < 100; i++ {
		var cell, _ = excelize.CoordinatesToCellName(1, i+1)
		err = ssw.SetRow(cell, []interface{}{i, i + 1, i + 2, i + 3}, true)
		if err != nil {
			log.Fatal(err)
		}

	}
	ssw.Flush()
	ef.Close()
}
