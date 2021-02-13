CREATE DATABASE IF NOT EXISTS tradesignals;
USE tradesignals;

create table TRADESIGNALS
(
    signalID               INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ClientID        TEXT,
    ClientName      TEXT,
    IGEpicID         TEXT,
    EpicName    TEXT,
    OpenOptionSpotPrice   DECIMAL(19,2),
    CloseOptionSpotPrice DECIMAL(19,2),
    OpenStockSpotPrice   DECIMAL(19,2),
    CloseStockSpotPrice   DECIMAL(19,2),
    StrikePrice   DECIMAL(19,2),
    OriginatingMessageID   TEXT,
    OriginatingMessageMedium   TEXT,
    TradeOptionDirection   TEXT,
    TradeOptionType   TEXT,
    SpeculativeTrade   BOOL,
    CoreTrade   BOOL,
    SignalPerformance   DECIMAL(19,2)
);