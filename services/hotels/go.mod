module hotels-service

go 1.23.4

require (
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/goli-nababa/golibaba-backend/common v0.0.0
	github.com/goli-nababa/golibaba-backend/modules/gateway_client v0.0.0
	github.com/goli-nababa/golibaba-backend/modules/user_service_client v0.0.0
	github.com/google/uuid v1.6.0
	github.com/jinzhu/copier v0.4.0
	go.uber.org/zap v1.27.0
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/goli-nababa/golibaba-backend/proto/pb v0.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
)

replace (
	github.com/goli-nababa/golibaba-backend/common => ../../common
	github.com/goli-nababa/golibaba-backend/modules/gateway_client => ../../modules/gateway_client
	github.com/goli-nababa/golibaba-backend/modules/user_service_client => ../../modules/user_service_client
	github.com/goli-nababa/golibaba-backend/proto/pb => ../../proto/pb
)
