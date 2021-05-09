package algorath

import (
	"algorath/endpoint"
	"algorath/manager"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/book"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/candle"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/event"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/status"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/ticker"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/trades"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/mux"
	"log"
	"algorath/repository"
)

func main() {

	database := repository.New()
	manager  := manager.New(database)

	endpoint.New(database, manager)

	m := mux.
		New().
		TransformRaw().
		WithAPIKEY("vTkXuCX73g1fG8xKiGNAXNAy5NdNcrRohzGoXzKecwI").
		WithAPISEC("weZNnklpR5X7U4LpwhlKCY38DgHBlxS073DhwGS5xRq").
		Start()


	for _, pair := range pairs {
		tradePld := event.Subscribe{
			Channel: "trades",
			Symbol:  "t" + pair,
		}

		tickPld := event.Subscribe{
			Channel: "ticker",
			Symbol:  "t" + pair,
		}

		candlesPld := event.Subscribe{
			Channel: "candles",
			Key:     "trade:1m:t" + pair,
		}

		rawBookPld := event.Subscribe{
			Channel:   "book",
			Precision: "R0",
			Symbol:    "t" + pair,
		}

		bookPld := event.Subscribe{
			Channel:   "book",
			Precision: "P0",
			Frequency: "F0",
			Symbol:    "t" + pair,
		}

		m.Subscribe(tradePld)
		m.Subscribe(tickPld)
		m.Subscribe(candlesPld)
		m.Subscribe(rawBookPld)
		m.Subscribe(bookPld)
	}

	derivStatusPld := event.Subscribe{
		Channel: "status",
		Key:     "deriv:tBTCF0:USTF0",
	}

	liqStatusPld := event.Subscribe{
		Channel: "status",
		Key:     "liq:global",
	}

	fundingPairTrade := event.Subscribe{
		Channel: "trades",
		Symbol:  "fUSD",
	}

	m.Subscribe(derivStatusPld)
	m.Subscribe(liqStatusPld)
	m.Subscribe(fundingPairTrade)

	crash := make(chan error)

	go func() {
		// if listener will fail, program will exit by passing error to crash channel
		crash <- m.Listen(func(msg interface{}, err error) {
			if err != nil {
				log.Printf("non crucial error received: %s\n", err)
			}

			switch v := msg.(type) {
			case event.Info:
				log.Printf("%T: %+v\n", v, v)
			case trades.TradeSnapshot:
				log.Printf("%T: %+v\n", v, v)
			case trades.FundingTradeSnapshot:
				log.Printf("%T: %+v\n", v, v)
			case trades.TradeExecutionUpdate:
				log.Printf("%T: %+v\n", v, v)
			case trades.TradeExecuted:
				log.Printf("%T: %+v\n", v, v)
			case trades.FundingTradeExecutionUpdate:
				log.Printf("%T: %+v\n", v, v)
			case trades.FundingTradeExecuted:
				log.Printf("%T: %+v\n", v, v)
			case *ticker.Ticker:
				log.Printf("%T: %+v\n", v, v)
			case *ticker.Snapshot:
				log.Printf("%T: %+v\n", v, v)
			case *book.Book:
				log.Printf("%T: %+v\n", v, v)
			case *book.Snapshot:
				log.Printf("%T: %+v\n", v, v)
			case *candle.Candle:
				log.Printf("%T: %+v\n", v, v)
			case *candle.Snapshot:
				log.Printf("%T: %+v\n", v, v)
			case *status.Derivative:
				log.Printf("%T: %+v\n", v, v)
			case *status.DerivativesSnapshot:
				log.Printf("%T: %+v\n", v, v)
			case *status.Liquidation:
				log.Printf("%T: %+v\n", v, v)
			case *status.LiquidationsSnapshot:
				log.Printf("%T: %+v\n", v, v)
			default:
				log.Printf("raw/unrecognized msg: %T: %s\n", v, v)
			}
		})
	}()

	log.Fatal(<-crash)
}