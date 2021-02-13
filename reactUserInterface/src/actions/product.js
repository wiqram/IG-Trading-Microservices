import axios from 'axios';
import * as constant from '../constants/product';

/** Get List **/
//const SERVICE_API = 'http://localhost:3070/products';
const SERVICE_GRPC_API = 'http://localhost:9090/v1/uibackend/getproducts';

/** When API is not available locally */
const GIT_MOCK_PRODUCTS_URL = 'https://raw.githubusercontent.com/nikusharma2010/demo/master/mobile-app/test-data/products.json';

const requestProducts = () => {
    return {
        type: constant.REQUEST_PRODUCTS
    };
};
const receiveProducts = (products) => {
    return {
        type: constant.RECEIVE_PRODUCTS,
        products: products.products
    };
};
const receiveError = (error) => {
    return {
        type: constant.RECEIVE_FAILURE,
        error: error
    };
};
/** return all Products list **/
const getProducts = () => {
    return async dispatch => {
        try {
            dispatch(requestProducts());
            console.log(">>>>>>>>>>>>>>we are in service API for products");
            //const res = await axios.get(SERVICE_API);
            //console.log(">>>>>>>>>>>>>>we are in service API for products")
            const res = await axios.get(SERVICE_GRPC_API);
            //console.log("this is the response from axios.get", res.data);
            dispatch(receiveProducts(res.data));
        } catch (error) {
          console.error(new Error('couldnt get results'));
          dispatch(receiveError(error));
        }
    }
};

export default getProducts;