package domain

import (
	"time"
)

type AnalyticsReport struct {
	ID          string
	BusinessID  uint64
	StartDate   time.Time
	EndDate     time.Time
	Metrics     map[string]float64
	Trends      map[string][]DataPoint
	Comparisons map[string]Comparison
	GeneratedAt time.Time
}

type DataPoint struct {
	Timestamp time.Time
	Value     float64
	Label     string
}

type Comparison struct {
	CurrentValue  float64
	PreviousValue float64
	ChangePercent float64
}

type KPIMetric struct {
	Name           string
	Value          float64
	Target         float64
	Unit           string
	TrendDirection string // "up", "down", "neutral"
	ComparedTo     time.Time
}

func CalculateGrowthRate(current, previous float64) float64 {
	if previous == 0 {
		return 0
	}
	return ((current - previous) / previous) * 100
}

func (ar *AnalyticsReport) AddMetric(name string, value float64) {
	if ar.Metrics == nil {
		ar.Metrics = make(map[string]float64)
	}
	ar.Metrics[name] = value
}

func (ar *AnalyticsReport) AddTrendPoint(name string, point DataPoint) {
	if ar.Trends == nil {
		ar.Trends = make(map[string][]DataPoint)
	}
	ar.Trends[name] = append(ar.Trends[name], point)
}
