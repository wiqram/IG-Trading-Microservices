import { combineReducers } from 'redux';
import auth from './auth';
import runtime from './runtime';
import navigation from './navigation';
import posts from './posts';
import productReducer from "./product";
import transactionReducer from "./transaction";
import tradeStatistics from "./tradeStatistics";

export default combineReducers({
  auth,
  runtime,
  navigation,
  posts,
  tradeStatistics,
  productState: productReducer,
  transactionState: transactionReducer,
});
