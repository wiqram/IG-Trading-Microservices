//const Koa = require('koa');
const Groa = require('groa');
const path = require('path');
const middleware = require('../middleware/middleware');
//const { productsRouter } = require('../products/routesGroa');
const PROTO_PATH = path.join(__dirname, '/../../proto/uibackend.proto');
/////////////////////////////////start
const Router = require('groa-router');
const productsRouter = new Router();
const cors = require('koa-cors');
const errorHandler = require('../middleware/error-handler');
const bodyParser = require('koa-bodyparser');
const Logger = require('koa-logger');
const { findProducts } = require('../products/service/find-products');
const { findTransactionsByProductId } = require('../products/service/find-transactions');
const messages = require('../../proto/uibackend_pb');
/*const { productMapper } = require('../products/mapper/mapper');
const NoDataFoundError = require('../error/NoDataFoundError');*/
const connectDB = require('../middleware/dbConnect');
const c2k = require('koa-connect');
const dbConnect = c2k(connectDB);
/////////////////////////////////end
const app = new Groa();
// Add proto file
app.addProto(PROTO_PATH);
//console.log(`proto path added to path for groa app -- `, PROTO_PATH);
/** Setup middleware */
//middleware(app);
/*app.use(cors());*/
app.use(dbConnect);
//app.use(bodyParser());
/*app.use(Logger());
app.use(errorHandler);*/

/** Setup Routers */

// package: demo_uibackend
// service: UibackendSvc
// method: getProducts
productsRouter.rpc('/demo_uibackend.UibackendSvc/GetProducts', async (ctx) => {
  //console.log('inside getProducts router method');
  let products = await findProducts(ctx);
  //console.log('we have recieved all products' + resp.getProductsList().length);
  ctx.body = { products: products };
  //const util = require('util');
  //let resp = ctx.res;
  //console.log('this is the response header' + resp.toString());
  //  console.log(util.inspect(resp, {depth: 2}));
});

// package: demo_uibackend
// service: UibackendSvc
// method: getTransactions
productsRouter.rpc('/demo_uibackend.UibackendSvc/GetTransactions', async (ctx) => {
  //console.log('inside getTransactions router method and this is the request');
  let transactions = await findTransactionsByProductId(ctx);
 // console.log('we have recieved all products' + resp.getTransactionsList().length);
  ctx.body = { transactions: transactions };
});

app.use(productsRouter.routes());

module.exports = app;


