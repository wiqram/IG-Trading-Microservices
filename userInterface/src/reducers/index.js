import { combineReducers } from 'redux';
import productReducer from '../bundle/product/middleware/reducer';
import transactionReducer from '../bundle/transaction/middleware/reducer';


const rootReducer = combineReducers(
    {
        productState: productReducer,
        transactionState: transactionReducer,
    }
);

export default rootReducer;
