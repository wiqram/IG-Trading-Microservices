import axios from 'axios';
import * as constant from './constants';
//const { Client } = require('groa');
//const path = require('path');
//const client = new Client('0.0.0.0', 3070);
//const PROTO_PATH = path.join(__dirname, '../../../../proto/uibackend.proto');
/** Get List **/
const SERVICE_API = 'http://localhost:3070/products';
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
const registerUser = () => {
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
// package: demo_uibackend
// service: UibackendSvc
// method: getProducts
/*const getProducts = async () => {
  console.log(">>>>>>>>>>>>>>we are in service API for products");
  dispatch(requestProducts());
  // Loading definition file
  await client.loadProto(PROTO_PATH);

  // Get service defnined
  let UibackendSvc = client.getService('demo_uibackend.UibackendSvc');
  console.log("this is the response from getservice", UibackendSvc.toString());
  // call
  let ret = await UibackendSvc.getProducts({

  });
  console.log(ret);
  dispatch(receiveProducts(ret));

};*/


export default registerUser;