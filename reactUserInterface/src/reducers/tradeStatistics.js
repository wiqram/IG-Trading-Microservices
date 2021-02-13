import {
  FETCH_ALL_TRADES_INITIAL,
  FETCH_ALL_TRADES_REQUEST,
  FETCH_ALL_TRADES_SUCCESS,
  FETCH_ALL_TRADES_FAILURE,
} from '../actions/tradeStatistics';

export default function tradeStatistics(
  state = {
    isFetching: false,
    trades:{},
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
        trades: action.trades,
      });
    case FETCH_ALL_TRADES_FAILURE:
      return Object.assign({}, state, {
        isFetching: false,
        message: 'Something wrong happened. Please come back later',
      });
    default:
      return state;
  }
}
