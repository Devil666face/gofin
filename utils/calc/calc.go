package calc

import (
	"math"

	"github.com/Devil666face/gofinabot/models"
)

var (
	BALANCE_KEY = "Балансовая"
	EXPENCE_KEY = "Расход"
	INCOME_KEY  = "Доход"
)

type ReportCalc struct {
	trans   []models.MoneyTransaction
	Sum     map[uint]int
	Percent map[uint]float64
	Expence int
	Income  int
	Balance int
}

func New(trans []models.MoneyTransaction) *ReportCalc {
	r := ReportCalc{}
	r.Sum = make(map[uint]int)
	r.Percent = make(map[uint]float64)
	r.trans = trans
	r.getTotals()
	r.getForTypes()
	r.getPercentsForTypes()
	return &r
}

func (r *ReportCalc) getTotals() {
	r.Expence = r.getExpence()
	r.Income = r.getIncome()
	r.Balance = r.Income + r.Expence
}

func (r *ReportCalc) getIncome() (income int) {
	for _, v := range r.trans {
		if v.Value > 0 {
			income += v.Value
		}
	}
	return
}

func (r *ReportCalc) getExpence() (expence int) {
	for _, v := range r.trans {
		if v.Value < 0 {
			expence += v.Value
		}
	}
	return

}

func (r *ReportCalc) getForTypes() {
	for _, v := range r.trans {
		if _, ok := r.Sum[v.TypeTransactionID]; !ok {
			r.Sum[v.TypeTransactionID] = v.Value
		} else {
			r.Sum[v.TypeTransactionID] += v.Value
		}
	}
}

func (r *ReportCalc) getPercentsForTypes() {
	for k, v := range r.Sum {
		if v > 0 {
			r.Percent[k] = 0
		} else {
			r.Percent[k] = math.Abs(float64(v) / float64(r.Income) * 100)
		}
	}
}
