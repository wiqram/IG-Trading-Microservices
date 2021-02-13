const { productMapper } = require('../mapper/mapper');
const NoDataFoundError = require('../../error/NoDataFoundError');
//const messages = require('../../../proto/uibackend_pb');

const findProducts = async (ctx) => {
  let products = [];
  let result = await ctx.req.dbConnect.query('SELECT * FROM PRODUCTS');
  if (result.length > 0) {
    products = productMapper(result);
    //console.log('this is the response in find products' + products.length)
    //return resp;
  } else {
    throw new NoDataFoundError();
  }
  return products;
};

module.exports = { findProducts };
