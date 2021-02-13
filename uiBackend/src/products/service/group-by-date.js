const groupByDate = (transactions) => Object.values(transactions.reduce((result, item) => {
  // Create new group
  if (!result[item.dateOn]) {
    result[item.dateOn] = {
      id: item.id,
      productId: item.productId,
      dateOn: item.dateOn,
      debitCredit: item.debitCredit,
      transDescription: item.transDescription,
      details: item.details,
      amount: item.amount,
      transactionList: []
    };
  }
  // Append to group
  result[item.dateOn].transactionList.push(item);
  return result;
}, {}));
module.exports = { groupByDate };

