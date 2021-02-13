const request = require('supertest')(process.env.SUT_URL);
const assert = require('assert');

describe('GET /products', () => {
  it('responds with content-type - json', (done) => {
    request
      .get('/products')
      .expect('Content-Type', /json/, done);
  });
  it('responds with status - 200', (done) => {
    request
      .get('/products')
      .set('Content-Type', 'application/json')
      .expect(200, done);
  });
  it('responds with products', (done) => {
    request
      .get('/products')
      .set('Content-Type', 'javascript/json')
      .expect(200)
      .end((err, res) => {
        if (err) return err;
        assert.strictEqual(res.body.products[0].productId, 'SA001');
        done();
      });
  });
});
