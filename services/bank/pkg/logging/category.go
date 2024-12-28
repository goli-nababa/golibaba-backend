package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	IO              Category = "IO"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Prometheus      Category = "Prometheus"
)

const (
	Startup             SubCategory = "Startup"
	ExternalService     SubCategory = "ExternalService"
	Migration           SubCategory = "Migration"
	Select              SubCategory = "Select"
	Rollback            SubCategory = "Rollback"
	Update              SubCategory = "Update"
	Delete              SubCategory = "Delete"
	Insert              SubCategory = "Insert"
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	FailedToCreateUser  SubCategory = "FailedToCreateUser"
	MobileValidation    SubCategory = "MobileValidation"
	PasswordValidation  SubCategory = "PasswordValidation"
	RemoveFile          SubCategory = "RemoveFile"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
	Duration     ExtraKey = "Duration"
	UserID       ExtraKey = "UserID"
)