package signalxtractor

// TradeSignal represents signal that client recommends. this is the structure on sql db plus there will date and timestamp for each signal
type TradeSignal struct {
	ClientID                 string  // id of the prof. trader recommending this trade signal
	ClientName               string  // name of the prof. trader recommending this trade signal
	SignalID                 string  // unique id generated automatically. this will allow us to identify how many traders hve played this trade
	IGEpicID                 string  // ID relating to the name as per IG broker
	EpicName                 string  // name from the message
	OpenOptionSpotPrice      string  //spot price of option when opened
	CloseOptionSpotPrice     float32 //spot price of option when closed
	OpenStockSpotPrice       string  //spot price of underlying asset when opened
	CloseStockSpotPrice      float32 //spot price of underlying asset when closed
	StrikePrice              float32 // recommended strike price to open
	OriginatingMessageID     string  // Unique ID from originator
	OriginatingMessageMedium string  // gmail or twitter
	TradeOptionDirection     string  // Buy or Sell. usually dont sell only buy put or call
	TradeOptionType          string  // Call or Put
	SpeculativeTrade         bool    //whether its spec or not
	CoreTrade                bool    // if its day/short term or long term trade. core trades are long term
	SignalPerformance        float32 //this is calculated as soon as position is closed based on open spot price and close spot price of the option bought
}
