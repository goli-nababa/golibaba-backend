package service

import (
	"bank_service/internal/services/financial_report/domain"
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

type exportHelper struct{}

func newExportHelper() exportHelper {
	return exportHelper{}
}

func (h exportHelper) exportToFormat(report *domain.FinancialReport, format domain.ReportFormat) ([]byte, error) {
	switch format {
	case domain.ReportFormatCSV:
		return h.exportToCSV(report)
	case domain.ReportFormatPDF:
		return nil, fmt.Errorf("PDF export not implemented")
	case domain.ReportFormatExcel:
		return nil, fmt.Errorf("Excel export not implemented")
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

func (h exportHelper) exportToCSV(report *domain.FinancialReport) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	headers := []string{"Metric", "Amount", "Currency"}
	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	for name, value := range report.Metrics {
		row := []string{
			name,
			strconv.FormatInt(value.Amount, 10),
			value.Currency,
		}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
