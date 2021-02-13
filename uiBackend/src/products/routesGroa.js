const Router = require('groa-router');
const productsRouter = new Router({
  //prefix: 'demo_uibackend.UibackendSvc'
});
const messages = require('../../proto/uibackend_pb');

const { findProducts } = require('./service/find-products');
const { findTransactionsByProductId } = require('./service/find-transactions');

const handlers = require('./handlers/handler');

// package: demo_uibackend
// service: UibackendSvc
// method: getProducts
productsRouter.rpc('/demo_uibackend.UibackendSvc/GetProducts', async (ctx) => {
  console.log('inside getProducts router method');
  //let resp = new messages.ProductsResponse();
  //resp.setProductsList(await findProducts(ctx));
  //const products = await findProducts(ctx);
  let resp = await findProducts(ctx);
  //resp.setProductsList(products);
  //console.log('we have recieved all products' + resp.getProductsList().length);
  ctx.body = resp;
});

// package: demo_uibackend
// service: UibackendSvc
// method: getTransactions
productsRouter.rpc('/demo_uibackend.UibackendSvc/GetTransactions', async (ctx) => {
  console.log('inside getTransactions router method and this is the request');
  let req = ctx.req.body;
  console.log('this is the ctx param - product ID - ' + req.productId);
  console.log('this is the ctx param - filterBy - ' + req.filterBy);
  //let resp = messages.TransactionsResponse();
  //resp.setTransactionsList(await findTransactionsByProductId(ctx));
  let resp = await findTransactionsByProductId(ctx);
  ctx.body = resp;
});
/*productsRouter.rpc('/demo_uibackend.UibackendSvc/getTransactions', handlers.getProducts);
productsRouter.rpc('/demo_uibackend.UibackendSvc/getTransactions', handlers.getTransactionsByProductId);*/

  /*productsRouter.get('/products', handlers.getProducts);
  productsRouter.get('/transactions/:productId', handlers.getTransactionsByProductId);*/

productsRouter.routes();
//productsRouter.allowedMethods();

module.exports = { productsRouter };
