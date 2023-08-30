package excel

import (
	"fmt"
	"log"
	"time"

	"github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils/calc"
	"github.com/xuri/excelize/v2"
)

var (
	DEFAULT_SHEET_NAME_FOR_DELETE = "Sheet1"
	DEFAULT_FILE_NAME             = "Finance report"
)

type ExcelReport struct {
	FileName   string
	SheetName  string
	trans      []models.MoneyTransaction
	file       *excelize.File
	calc       *calc.ReportCalc
	sheetIndex int
}

func New(trans []models.MoneyTransaction) *ExcelReport {
	r := ExcelReport{}
	r.trans = trans
	r.FileName = r.getFileName()
	r.SheetName = r.getSheetName()
	if err := r.NewFile(); err != nil {
		log.Print(err)
	}
	r.calc = calc.New(trans)
	r.setCellValues()
	r.setTotalsCellValues()
	return &r
}

func (r *ExcelReport) NewFile() error {
	r.file = excelize.NewFile()
	index, err := r.file.NewSheet(r.SheetName)
	if err != nil {
		return err
	}
	r.sheetIndex = index
	if err := r.file.DeleteSheet(DEFAULT_SHEET_NAME_FOR_DELETE); err != nil {
		log.Print(err)
	}
	return nil
}

func (r *ExcelReport) Save() error {
	if err := r.file.SaveAs(r.FileName); err != nil {
		return err
	}
	if err := r.file.Close(); err != nil {
		return err
	}
	return nil
}

func (r *ExcelReport) setCellValues() {
	r.file.SetActiveSheet(r.sheetIndex)
	for i, v := range r.trans {
		r.set(1, i+1, fmt.Sprintf("%s", v.CreatedAt))
		r.set(2, i+1, fmt.Sprintf("%s", v.TypeTransaction()))
		r.set(3, i+1, fmt.Sprintf("%s", v.Balance()))
		r.set(4, i+1, v.Comment)
		r.set(5, i+1, v.Value)
	}
}

func (r *ExcelReport) setTotalsCellValues() {
	i := len(r.trans) + 2
	for k, v := range r.calc.Sum {
		trtype := models.TypeTransaction{}
		if err := trtype.Get(k); err != nil {
			log.Print(err)
		}
		r.set(3, i, fmt.Sprintf("%s", trtype))
		r.set(4, i, v)
		r.set(5, i, fmt.Sprintf("%.2f", r.calc.Percent[k]))
		i++
	}
	r.set(3, i+1, calc.BALANCE_KEY)
	r.set(3, i+2, calc.INCOME_KEY)
	r.set(3, i+3, calc.EXPENCE_KEY)
	r.set(4, i+1, r.calc.Balance)
	r.set(4, i+2, r.calc.Income)
	r.set(4, i+3, r.calc.Expence)
}

func (r *ExcelReport) set(cell, row int, value interface{}) {
	if err := r.file.SetCellValue(r.SheetName, cellIndex(cell, row), value); err != nil {
		log.Print(err)
	}
}

func cellIndex(col, row int) string {
	if cell, err := excelize.CoordinatesToCellName(col, row); err == nil {
		return cell
	}
	return ""
}

func (r ExcelReport) getFileName() string {
	return fmt.Sprintf("%s %02d.%02d.xlsx", DEFAULT_FILE_NAME, time.Now().Month(), time.Now().Year())
}

func (r ExcelReport) getSheetName() string {
	return fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
}
