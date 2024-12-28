module github.com/goli-nababa/golibaba-backend/modules/bank_service_client

go 1.23.3

require (
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.10.0
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.36.0
	github.com/goli-nababa/golibaba-backend/proto/pb v0.0.0

)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

)





replace github.com/goli-nababa/golibaba-backend/proto/pb => ../../proto/pb