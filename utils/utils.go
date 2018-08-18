package utils

import (
	"fmt"
	"math"
	"time"
)

const YearMonths = 12.0

var counter = 0

type PlanDetail struct {
	BorrowerPaymentAmount         float64 `json:"BorrowerPaymentAmount"`
	Date                          string  `json:"Date"`
	InitialOutstandingPrincipal   float64 `json:"InitialOutstandingPrincipal"`
	Interest                      float64 `json:"Interest"`
	Principal                     float64 `json:"Principal"`
	RemainingOutstandingPrincipal float64 `json:"RemainingOutstandingPrincipal"`
}
type PlanDetailList struct {
	List []PlanDetail
}

/**
* AddItem
* To add PlanDetail struct to list
 */
func (plan_detail_list *PlanDetailList) AddItem(item PlanDetail) []PlanDetail {
	plan_detail_list.List = append(plan_detail_list.List, item)
	return plan_detail_list.List
}
func (p *PlanDetailList) GetList() []PlanDetail {
	return p.List
}

/**
* Annuity
* calculates annuity
 */
func Annuity(rate float64, principle float64, year float64) float64 {
	if rate == 0 {
		return 0
	}
	if principle == 0 {
		return 0
	}
	if year == 0 {
		return 0
	}
	ratePerMonth := ConvertRateToMonthly(rate)
	annuity := (principle * ratePerMonth) * (math.Pow((1 + ratePerMonth), (year * YearMonths))) / ((math.Pow(1+ratePerMonth, year*YearMonths)) - 1)
	return annuity
}

/**
* CalculateInterest
* CalculateInterest the interest for each amount
 */
func CalculateInterest(rate float64, principle float64) float64 {
	if rate == 0 {
		return 0
	}
	if principle == 0 {
		return 0
	}
	ratePerMonth := ConvertRateToMonthly(rate)
	monthInterest := ratePerMonth * principle
	return monthInterest
}

/**
* ConvertRateToMonthly
* Convert rate to monthly
 */
func ConvertRateToMonthly(rate float64) float64 {
	return (rate / 100) / 12
}

var items = PlanDetailList{}

/**
* GenerateRepaymentPlan
* function that generate loan plan recurssively
 */
func GenerateRepaymentPlan(rate float64, amount float64, year float64, startDate string) []PlanDetail {
	if rate == 0 || amount == 0 || year == 0 {
		return PlanDetailList{}.List
	}
	var planDetail = PlanDetail{}
	date, _ := time.Parse(time.RFC3339, startDate)
	date = date.AddDate(0, counter, 0)
	if counter < 1 {
		planDetail.BorrowerPaymentAmount = Annuity(rate, amount, year)
	} else {

		planDetail.BorrowerPaymentAmount = items.List[0].BorrowerPaymentAmount
	}
	planDetail.Date = date.String()
	planDetail.InitialOutstandingPrincipal = amount

	planDetail.Interest = CalculateInterest(rate, amount)
	planDetail.Principal = planDetail.BorrowerPaymentAmount - planDetail.Interest

	planDetail.RemainingOutstandingPrincipal = planDetail.InitialOutstandingPrincipal - planDetail.Principal
	if planDetail.RemainingOutstandingPrincipal < planDetail.BorrowerPaymentAmount {
	}
	items.AddItem(planDetail)
	counter = counter + 1
	fmt.Println(planDetail)
	if planDetail.RemainingOutstandingPrincipal < planDetail.BorrowerPaymentAmount {
		return items.GetList()
	}
	GenerateRepaymentPlan(rate, planDetail.RemainingOutstandingPrincipal, year, startDate)

	return items.GetList()
}
