const Router = require('koa-router');
const productsRouter = new Router();

const handlers = require('./handlers/handler');

productsRouter.get('/products', handlers.getProducts);
productsRouter.get('/transactions/:productId', handlers.getTransactionsByProductId);

productsRouter.routes();
productsRouter.allowedMethods();

module.exports = { productsRouter };
