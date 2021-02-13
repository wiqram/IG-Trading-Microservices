const { findProducts } = require('../service/find-products');
const { findTransactionsByProductId } = require('../service/find-transactions');

/**
 * main handler for the finder
 */
const handlers = {
  getProducts: async ctx => {
    /** * implement logic */
    const products = await findProducts(ctx);
    /**  return body and status code */
    ctx.body = {
      products
    };
  },
  getTransactionsByProductId: async ctx => {
    /** * implement logic */
    const transactions = await findTransactionsByProductId(ctx);
    /**  return body and status code */
    ctx.body = {
      transactions
    };
  }
};

module.exports = handlers;
