import * as constant from './constants';

const initialState = {
    classes: {},
    value: 0
};
const registerReducer = (state = initialState, action) => {
    switch (action.type) {
    case constant.REQUEST_PRODUCTS:
    {
        console.log("this is inside requestproduct")
        return { ...state, isLoading: true };
    }

    case constant.RECEIVE_FAILURE:
    {
        return { ...state, isLoading: false, error: action.error, };
    }

    case constant.RECEIVE_PRODUCTS:
    {
        return { ...state, isLoading: false, products: action.products, };
    }

    default:
        return state;
    }
};
export default registerReducer;