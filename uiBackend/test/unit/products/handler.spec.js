const assert = require('assert');
const products = require('./data/products');
const transactionsAll = require('./data/transactions-all');
const transactionsGroup = require('./data/transactions-group-all');
const transactionsCr = require('./data/transactions-cr');
const transactionsGroupCr = require('./data/transactions-group-cr');
const transactionsDr = require('./data/transactions-dr');
const transactionsGroupDr = require('./data/transactions-group-dr');
describe('Handler', () => {
  const execute = () => {
    const sut = require('../../../src/products/handlers/handler');
    return sut;
  };
  describe('Find all products - Success', () => {
    it('should load all list of products', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => products
          }
        },
        body: {
        }
      };

      await execute().getProducts(ctx);

      assert.deepStrictEqual(ctx.body.products, products);
    });
  });
  describe('Find all products - Error', () => {
    it('should throw error', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => []
          }
        },
        body: {
        }
      };

      const NoDataFoundError =
      {
        title: 'No Data Found',
        message: 'No such data found, try again !!!',
        statusCode: 422
      };

      try {
        await execute().getProducts(ctx);
      } catch (error) {
        assert.deepStrictEqual(error.statusCode, NoDataFoundError.statusCode);
      }
    });
  });

  describe('getTransactionsByProductId - ALL ', () => {
    it('should load all list of trans', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => transactionsAll
          }
        },
        body: {
        },
        params: { productId: 'SA001' },
        query: { filterBy: null }
      };

      await execute().getTransactionsByProductId(ctx);

      assert.deepStrictEqual(ctx.body.transactions, transactionsGroup);
    });
  });

  describe('getTransactionsByProductId - only CR', () => {
    it('should load all list of trans', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => transactionsCr
          }
        },
        body: {
        },
        params: { productId: 'SA001' },
        query: { filterBy: 'CR' }
      };

      await execute().getTransactionsByProductId(ctx);

      assert.deepStrictEqual(ctx.body.transactions, transactionsGroupCr);
    });
  });

  describe('getTransactionsByProductId - only DR', () => {
    it('should load all list of trans', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => transactionsDr
          }
        },
        body: {
        },
        params: { productId: 'SA001' },
        query: { filterBy: 'DR' }
      };

      await execute().getTransactionsByProductId(ctx);

      assert.deepStrictEqual(ctx.body.transactions, transactionsGroupDr);
    });
  });
  describe('getTransactionsByProductId - Error', () => {
    it('should throw error', async () => {
      const ctx = {
        req: {
          dbConnect: {
            query: async (str) => []
          }
        },
        body: {
        },
        params: { productId: 'SA001' },
        query: { filterBy: 'DR' }
      };

      const NoDataFoundError =
      {
        title: 'No Data Found',
        message: 'No such data found, try again !!!',
        statusCode: 422
      };

      try {
        await execute().getTransactionsByProductId(ctx);
      } catch (error) {
        assert.deepStrictEqual(error.statusCode, NoDataFoundError.statusCode);
      }
    });
  });
});
