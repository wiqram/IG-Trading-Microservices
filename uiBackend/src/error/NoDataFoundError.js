
class NoDataFoundError {
  constructor (message, statusCode) {
    this.title = 'No Data Found';
    this.message = message || 'No such data found, try again !!!';
    this.statusCode = statusCode || 422;
  }
}

module.exports = NoDataFoundError;
