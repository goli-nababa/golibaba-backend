module github.com/goli-nababa/golibaba-backend/modules/user_service_client

go 1.23.3

require (
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.69.2
)

require google.golang.org/protobuf v1.36.0 // indirect

require (
	github.com/goli-nababa/golibaba-backend/proto/pb v0.0.0
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)

replace github.com/goli-nababa/golibaba-backend/proto/pb => ../../proto/pb
