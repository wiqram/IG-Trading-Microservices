import * as constant from './constants';

const initialState = {
    isLoading: false,
    error: '',
    transactions: []
};
const transactionReducer = (state = initialState, action) => {
    switch (action.type) {
    case constant.REQUEST_TRANSACTIONS:
    {
        return { ...state, isLoading: true, error: '', transactions: [] };
    }
    case constant.RECEIVE_FAILURE:
    {
        return { ...state, isLoading: false, error: action.error, transactions: [] };
    }
    case constant.RECEIVE_TRANSACTIONS:
    {
        return { ...state, isLoading: false, error: '', transactions: action.transactions };
    }

    default:
        return state;
    }
};
export default transactionReducer;