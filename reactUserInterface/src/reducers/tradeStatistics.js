import {
  FETCH_ALL_TRADES_INITIAL,
  FETCH_ALL_TRADES_REQUEST,
  FETCH_ALL_TRADES_SUCCESS,
  FETCH_ALL_TRADES_FAILURE,
  FETCH_MARKETSEARCH_FAILURE,
  FETCH_MARKETSEARCH_REQUEST,
  FETCH_MARKETSEARCH_SUCCESS
} from '../actions/tradeStatistics';

export default function tradeStatistics(
  state = {
    isFetching: false,
    tradeSentiment:{},
    searchTerm:"",
    markets: {},
    marketSearchNames:{},
    },
  action,
) {
  switch (action.type) {
    case FETCH_ALL_TRADES_INITIAL:
      return Object.assign({}, state, {
        isFetching: false,
      });
    case FETCH_ALL_TRADES_REQUEST:
      return Object.assign({}, state, {
        isFetching: true,
      });
    case FETCH_ALL_TRADES_SUCCESS:
      return Object.assign({}, state, {
        isFetching: false,
        tradeSentiment: action.tradeSentiment,
      });
    case FETCH_ALL_TRADES_FAILURE:
      return Object.assign({}, state, {
        isFetching: false,
        message: 'Something wrong happened. Please come back later',
      });
    case FETCH_MARKETSEARCH_REQUEST:
      return Object.assign({}, state, {
        isFetching: true,
      });
    case FETCH_MARKETSEARCH_SUCCESS:
      return Object.assign({}, state, {
        isFetching: false,
        markets: action.markets,
        marketSearchNames: action.marketSearchNames,
      });
    case FETCH_MARKETSEARCH_FAILURE:
      return Object.assign({}, state, {
        isFetching: false,
        message: 'Something wrong happened. Please come back later',
      });
    default:
      return state;
  }
}
