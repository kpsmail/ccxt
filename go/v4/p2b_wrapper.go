package ccxt

type P2b struct {
   *p2b
   Core *p2b
}

func NewP2b(userConfig map[string]interface{}) P2b {
   p := &p2b{}
   p.Init(userConfig)
   return P2b{
       p2b: p,
       Core:  p,
   }
}

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


/**
 * @method
 * @name p2b#fetchMarkets
 * @description retrieves data on all markets for bigone
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#markets
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} an array of objects representing market data
 */
func (this *P2b) FetchMarkets(params ...interface{}) ([]MarketInterface, error) {
    res := <- this.Core.FetchMarkets(params...)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewMarketInterfaceArray(res), nil
}
/**
 * @method
 * @name p2b#fetchTickers
 * @description fetches price tickers for multiple markets, statistical information calculated over the past 24 hours for each market
 * @see https://futures-docs.poloniex.com/#get-real-time-ticker-of-all-symbols
 * @param {string[]|undefined} symbols unified symbols of the markets to fetch the ticker for, all market tickers are returned if not assigned
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a dictionary of [ticker structures]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *P2b) FetchTickers(options ...FetchTickersOptions) (Tickers, error) {

    opts := FetchTickersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbols interface{} = nil
    if opts.Symbols != nil {
        symbols = *opts.Symbols
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTickers(symbols, params)
    if IsError(res) {
        return Tickers{}, CreateReturnError(res)
    }
    return NewTickers(res), nil
}
/**
 * @method
 * @name p2b#fetchTicker
 * @description fetches a price ticker, a statistical calculation with the information calculated over the past 24 hours for a specific market
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#ticker
 * @param {string} symbol unified symbol of the market to fetch the ticker for
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [ticker structure]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *P2b) FetchTicker(symbol string, options ...FetchTickerOptions) (Ticker, error) {

    opts := FetchTickerOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTicker(symbol, params)
    if IsError(res) {
        return Ticker{}, CreateReturnError(res)
    }
    return NewTicker(res), nil
}
/**
 * @method
 * @name p2b#fetchOrderBook
 * @description fetches information on open orders with bid (buy) and ask (sell) prices, volumes and other data
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#depth-result
 * @param {string} symbol unified symbol of the market to fetch the order book for
 * @param {int} [limit] the maximum amount of order book entries to return
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 *
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {string} [params.interval] 0 (default), 0.00000001, 0.0000001, 0.000001, 0.00001, 0.0001, 0.001, 0.01, 0.1, 1
 * @returns {object} A dictionary of [order book structures]{@link https://docs.ccxt.com/#/?id=order-book-structure} indexed by market symbols
 */
func (this *P2b) FetchOrderBook(symbol string, options ...FetchOrderBookOptions) (OrderBook, error) {

    opts := FetchOrderBookOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrderBook(symbol, limit, params)
    if IsError(res) {
        return OrderBook{}, CreateReturnError(res)
    }
    return NewOrderBook(res), nil
}
/**
 * @method
 * @name p2b#fetchTrades
 * @description get the list of most recent trades for a particular symbol
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#history
 * @param {string} symbol unified symbol of the market to fetch trades for
 * @param {int} [since] timestamp in ms of the earliest trade to fetch
 * @param {int} [limit] 1-100, default=50
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} params.lastId order id
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=public-trades}
 */
func (this *P2b) FetchTrades(symbol string, options ...FetchTradesOptions) ([]Trade, error) {

    opts := FetchTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name p2b#fetchOHLCV
 * @description fetches historical candlestick data containing the open, high, low, and close price, and the volume of a market
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#kline
 * @param {string} symbol unified symbol of the market to fetch OHLCV data for
 * @param {string} timeframe 1m, 1h, or 1d
 * @param {int} [since] timestamp in ms of the earliest candle to fetch
 * @param {int} [limit] 1-500, default=50
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} [params.offset] default=0, with this value the last candles are returned
 * @returns {int[][]} A list of candles ordered as timestamp, open, high, low, close, volume
 */
func (this *P2b) FetchOHLCV(symbol string, options ...FetchOHLCVOptions) ([]OHLCV, error) {

    opts := FetchOHLCVOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var timeframe interface{} = nil
    if opts.Timeframe != nil {
        timeframe = *opts.Timeframe
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOHLCV(symbol, timeframe, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOHLCVArray(res), nil
}
/**
 * @method
 * @name p2b#fetchBalance
 * @description query for balance and get the amount of funds available for trading or funds locked in orders
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#all-balances
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [balance structure]{@link https://docs.ccxt.com/#/?id=balance-structure}
 */
func (this *P2b) FetchBalance(params ...interface{}) (Balances, error) {
    res := <- this.Core.FetchBalance(params...)
    if IsError(res) {
        return Balances{}, CreateReturnError(res)
    }
    return NewBalances(res), nil
}
/**
 * @method
 * @name p2b#createOrder
 * @description create a trade order
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#create-order
 * @param {string} symbol unified symbol of the market to create an order in
 * @param {string} type must be 'limit'
 * @param {string} side 'buy' or 'sell'
 * @param {float} amount how much of currency you want to trade in units of base currency
 * @param {float} price the price at which the order is to be fulfilled, in units of the quote currency
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} an [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *P2b) CreateOrder(symbol string, typeVar string, side string, amount float64, options ...CreateOrderOptions) (Order, error) {

    opts := CreateOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var price interface{} = nil
    if opts.Price != nil {
        price = *opts.Price
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CreateOrder(symbol, typeVar, side, amount, price, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name p2b#cancelOrder
 * @description cancels an open order
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#cancel-order
 * @param {string} id order id
 * @param {string} symbol unified symbol of the market the order was made in
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} An [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *P2b) CancelOrder(id string, options ...CancelOrderOptions) (Order, error) {

    opts := CancelOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CancelOrder(id, symbol, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name p2b#fetchOpenOrders
 * @description fetch all unfilled currently open orders
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#open-orders
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for
 * @param {int} [limit] the maximum number of order structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 *
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {int} [params.offset] 0-10000, default=0
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *P2b) FetchOpenOrders(options ...FetchOpenOrdersOptions) ([]Order, error) {

    opts := FetchOpenOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOpenOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name p2b#fetchOrderTrades
 * @description fetch all the trades made from a single order
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#deals-by-order-id
 * @param {string} id order id
 * @param {string} symbol unified market symbol
 * @param {int} [since] the earliest time in ms to fetch trades for
 * @param {int} [limit] 1-100, default=50
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 *
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {int} [params.offset] 0-10000, default=0
 * @returns {object[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=trade-structure}
 */
func (this *P2b) FetchOrderTrades(id string, options ...FetchOrderTradesOptions) ([]Trade, error) {

    opts := FetchOrderTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrderTrades(id, symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name p2b#fetchMyTrades
 * @description fetch all trades made by the user, only the transaction records in the past 3 month can be queried, the time between since and params["until"] cannot be longer than 24 hours
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#deals-history-by-market
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for, default = params["until"] - 86400000
 * @param {int} [limit] 1-100, default=50
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} [params.until] the latest time in ms to fetch orders for, default = current timestamp or since + 86400000
 *
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {int} [params.offset] 0-10000, default=0
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=public-trades}
 */
func (this *P2b) FetchMyTrades(options ...FetchMyTradesOptions) ([]Trade, error) {

    opts := FetchMyTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchMyTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name p2b#fetchClosedOrders
 * @description fetches information on multiple closed orders made by the user, the time between since and params["untnil"] cannot be longer than 24 hours
 * @see https://github.com/P2B-team/p2b-api-docs/blob/master/api-doc.md#orders-history-by-market
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for, default = params["until"] - 86400000
 * @param {int} [limit] 1-100, default=50
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} [params.until] the latest time in ms to fetch orders for, default = current timestamp or since + 86400000
 *
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {int} [params.offset] 0-10000, default=0
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *P2b) FetchClosedOrders(options ...FetchClosedOrdersOptions) ([]Order, error) {

    opts := FetchClosedOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchClosedOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}