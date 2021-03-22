module github.com/bitfinexcom/bfxfixgw

go 1.15

require (
	github.com/bitfinexcom/bitfinex-api-go v0.0.0-20210226135753-3649a968eee2
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/quickfixgo/enum v0.0.0-20171007195659-2cbed3730c3e
	github.com/quickfixgo/field v0.0.0-20171007195410-74cea5ec78c7
	github.com/quickfixgo/fix42 v0.0.0-20171007212724-86a4567f1c77
	github.com/quickfixgo/fix44 v0.0.0-20171007213039-f090a1006218
	github.com/quickfixgo/fix50 v0.0.0-20171007213247-d09e70735b64
	github.com/quickfixgo/fixt11 v0.0.0-20171007213433-d9788ca97f5d
	github.com/quickfixgo/quickfix v0.6.1-0.20190117144057-5bcdd16d49b0
	github.com/quickfixgo/tag v0.0.0-20171007194743-cbb465760521
	github.com/satori/go.uuid v1.1.1-0.20170321230731-5bf94b69c6b6
	github.com/shopspring/decimal v1.1.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.6.1
	go.uber.org/atomic v1.2.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.7.2-0.20171031232209-f85c78b1dd99
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/bitfinexcom/bitfinex-api-go => ../bitfinex-api-go
