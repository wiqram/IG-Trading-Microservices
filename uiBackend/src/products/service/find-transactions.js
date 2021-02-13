const { transactionMapper } = require('../mapper/mapper');
const NoDataFoundError = require('../../error/NoDataFoundError');
const { groupByDate } = require('./group-by-date');
//const messages = require('../../../proto/uibackend_pb');

const findTransactionsByProductId = async (ctx) => {
  let transactions = [];
  let req = ctx.req.body;
  let productId = req.productId;
  let filterBy = req.filterBy;
  let SQL;
  if (filterBy && filterBy!='ALL') {
    SQL = 'SELECT * FROM TRANSACTIONS where productId ="' + productId + '" AND debitCredit IN ("' + filterBy + '")';
  } else {
    SQL = 'SELECT * FROM TRANSACTIONS where productId ="' + productId + '"';
  }
  //console.log('this is the SQL statement' + SQL.toString());
  let result = await ctx.req.dbConnect.query(SQL);
  //console.log('this is the response right now before sending it back query ---- ' + result.toString());
  transactions = transactionMapper(result);
  //console.log('this is the response right now before sending it back ---- ' + resp.toString());
  if (transactions.length <= 0) {
    throw new NoDataFoundError();
  }
  /** Sort the transactions first  */
  /** * Group them in order  */
  const transactionsGrouped = groupByDate(transactions);
 // console.log('this is the response of trnasctions - ' + resp.getTransactionsList().length);
/*
  const util = require('util');
  console.log('this is the response header >>>>>>>>>>>>>>>>>' + transactionsGrouped.length);
  console.log(util.inspect(transactionsGrouped, {depth: 4}));
*/
  return transactionsGrouped;
};

module.exports = { findTransactionsByProductId };
