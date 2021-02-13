import * as constant from '../constants/transaction';
import axios from 'axios';
/** Get List **/
const SERVICE_API = 'http://localhost:3070/transactions/';
const SERVICE_GRPC_API = 'http://0.0.0.0:9090/v1/uibackend/gettransactions';

/** When API is not available locally */
const GIT_MOCK_ALL_TR_URL = 'https://raw.githubusercontent.com/nikusharma2010/demo/tree/master/mobile-app/test-data/transactions.json';
const GIT_MOCK_CR_TR_URL = 'https://raw.githubusercontent.com/nikusharma2010/demo/tree/master/mobile-app/test-data/transactions-cr.json';
const GIT_MOCK_DR_TR_URL = 'https://raw.githubusercontent.com/nikusharma2010/demo/tree/master/mobile-app/test-data/transactions-dr.json';


const requestTransactions = () => {
    return {
        type: constant.REQUEST_TRANSACTIONS
    };
};
const receiveTransactions = (transactions) => {
    return {
        type: constant.RECEIVE_TRANSACTIONS,
        transactions: transactions.transactions
    };
};
const receiveError = (error) => {
    return {
        type: constant.RECEIVE_FAILURE,
        error: error
    };
};
/** return all Transactions list **/
export const getTransactions = (productId, filterBy) => {
    return async dispatch => {
        try {
            dispatch(requestTransactions());
            console.log('we are ui get transactions function')
            const url = SERVICE_GRPC_API+"/"+productId+"/"+filterBy;
            console.log("this is the url " + url);
            const res = await axios.get(url);
            console.log("this is response" + res.data.length);
            /*let res;
            if (filterBy === 'CR') {
                res = await axios.get(GIT_MOCK_CR_TR_URL);
            } else if (filterBy === 'DR') {
                res = await axios.get(GIT_MOCK_DR_TR_URL);
            } else {
                res = await axios.get(GIT_MOCK_ALL_TR_URL);
            }*/
            dispatch(receiveTransactions(res.data));
        } catch (error) {
            dispatch(receiveError(error))
        }
    }
};
