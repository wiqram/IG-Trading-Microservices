import configureMockStore from 'redux-mock-store'
import thunk from 'redux-thunk'
import moxios from 'moxios';
import getProducts from '../../../../../src/bundle/product/middleware/actions'
import * as types from '../../../../../src/bundle/product/middleware/constants'

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)

describe('async products actions', () => {

    const products = [{
        productName: 'Saving A/C',
        currentBalance: '11.11',
        accountNumber: '111111',
        sortCode: '121212',
        availableBalance: '1212.12'
    }]

    beforeEach(() => {
        moxios.install();
    })

    afterEach(function () {
        moxios.uninstall();
    });
    it('creates RECEIVE_PRODUCTS when fetching products has been done', async () => {
        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.respondWith({
                status: 200,
                response: { products: products },
            });
        });
        const expectedActions = [
            { type: types.REQUEST_PRODUCTS },
            { type: types.RECEIVE_PRODUCTS, products: products }
        ]
        const store = mockStore({
            products: []
        })

        return store.dispatch(getProducts()).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
    it('creates RECEIVE_FAILURE when fetching products has been done', async () => {
        const errorResp = 'invalid data';

        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.reject(errorResp);
        });

        const expectedActions = [
            { type: types.REQUEST_PRODUCTS },
            { type: types.RECEIVE_FAILURE, error: errorResp }
        ]
        const store = mockStore({
            error: 'Error Occurred',
        })

        return store.dispatch(getProducts()).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
})