module transportation

go 1.23.3

toolchain go1.23.4

require (
	github.com/goli-nababa/golibaba-backend/common v0.0.0-00010101000000-000000000000
	github.com/goli-nababa/golibaba-backend/modules/gateway_client v0.0.0-00010101000000-000000000000
	github.com/goli-nababa/golibaba-backend/proto/pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.69.2
	gorm.io/gorm v1.25.12
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
)

require (
	github.com/google/uuid v1.6.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/rs/zerolog v1.33.0
	github.com/streadway/amqp v1.1.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20241215155358-4a5509556b9e
	golang.org/x/text v0.19.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gorm.io/driver/postgres v1.5.11
)

replace (
	github.com/goli-nababa/golibaba-backend/common => ../../common
	github.com/goli-nababa/golibaba-backend/modules/gateway_client => ../../modules/gateway_client
	github.com/goli-nababa/golibaba-backend/proto/pb => ../../proto/pb
)
