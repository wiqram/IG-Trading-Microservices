const request = require('supertest')(process.env.SUT_URL);
const assert = require('assert');

const URL = '/transactions/SA001';
const URL_EXP = '/transactions/SA0012';

describe('GET /transactions/:productId', () => {
  it('responds with content-type - json', (done) => {
    request
      .get(URL)
      .expect('Content-Type', /json/, done);
  });
  it('responds with status - 200', (done) => {
    request
      .get(URL)
      .set('Content-Type', 'application/json')
      .expect(200, done);
  });
  it('responds with transaction', (done) => {
    request
      .get(URL)
      .set('Content-Type', 'javascript/json')
      .expect(200)
      .end((err, res) => {
        if (err) return err;
        assert.strictEqual(res.body.transactions[0].transactionList[0].productId, 'SA001');
        done();
      });
  });
  it('responds with error - code 422', (done) => {
    request
      .get(URL_EXP)
      .set('Content-Type', 'application/json')
      .expect(422, done);
  });
  it('responds with error', (done) => {
    request
      .get(URL_EXP)
      .set('Content-Type', 'application/json')
      .expect(422)
      .end((err, res) => {
        assert.strictEqual(res.body.message, 'No such data found, try again !!!');
        assert.strictEqual(res.body.statusCode, 422);
        if (err) return err;
        done();
      });
  });
});
