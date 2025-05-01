module vehicles

go 1.23.3

toolchain go1.23.4

require (
	github.com/goli-nababa/golibaba-backend/modules/trip_service_client v0.0.0-00010101000000-000000000000
	github.com/goli-nababa/golibaba-backend/proto/pb v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.12
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/streadway/amqp v1.1.0
	golang.org/x/text v0.19.0 // indirect
	gorm.io/driver/postgres v1.5.11
)

replace (
	github.com/goli-nababa/golibaba-backend/common => ../../common
	github.com/goli-nababa/golibaba-backend/modules/gateway_client => ../../modules/gateway_client
	github.com/goli-nababa/golibaba-backend/modules/trip_service_client => ../../modules/trip_service_client
	github.com/goli-nababa/golibaba-backend/proto/pb => ../../proto/pb
)
