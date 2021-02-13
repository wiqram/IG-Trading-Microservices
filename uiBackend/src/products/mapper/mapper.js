/** * Map the data for JSON */
const productMapper = (data) => {
  let products = [];
  // if (!data) return null;

  data.forEach(value => {
    //console.log('this is the productId in mapper -- ' + value.productId);

    products.push({
      'id': value.id,
      'productId': value.productId,
      'productName': value.productName,
      'sortCode': value.sortCode,
      'accountNumber': value.accountNumber,
      'currentBalance': value.currentBalance,
      'availableBalance': value.availableBalance
    });
  });
  return products;
};

const transactionMapper = (data) => {
  let transactions = [];
  // if (!data) return null;

  data.forEach(value => {
    transactions.push({
      'id': value.id,
      'productId': value.productId,
      'dateOn': value.dateOn,
      'debitCredit': value.debitCredit,
      'transDescription': value.transDescription,
      'details': value.details,
      'amount': value.amount
    });
  });
  return transactions;
};
module.exports = { productMapper, transactionMapper };
