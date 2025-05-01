package domain

type ReportFormat string

const (
	ReportFormatCSV   ReportFormat = "csv"
	ReportFormatPDF   ReportFormat = "pdf"
	ReportFormatExcel ReportFormat = "excel"
)
