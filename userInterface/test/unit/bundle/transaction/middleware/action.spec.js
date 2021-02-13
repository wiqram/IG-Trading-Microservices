import configureMockStore from 'redux-mock-store'
import thunk from 'redux-thunk'
import moxios from 'moxios';
import * as actions from '../../../../../src/bundle/transaction/middleware/actions'
import * as types from '../../../../../src/bundle/transaction/middleware/constants'

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)

describe('async transaction actions', () => {

    const transactions = [
        {
            'id': 1,
            'productId': 'SA001',
            'dateOn': '01-01-2019',
            'debitCredit': 'CR',
            'transDescription': 'ASDA Shopping',
            'details': 'Trans at ADSA Livingston',
            'amount': 10010.2
        },
        {
            'id': 2,
            'productId': 'SA001',
            'dateOn': '02-01-2019',
            'debitCredit': 'DR',
            'transDescription': 'MORRISON Shopping',
            'details': 'Trans at MORRISON Livingston',
            'amount': 10020.2
        }];

    beforeEach(() => {
        moxios.install();
    })

    afterEach(function () {
        moxios.uninstall();
    });
    it('creates RECEIVE_TRANSACTIONS when fetching ALL transactions has been done', async () => {
        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.respondWith({
                status: 200,
                response: { transactions: transactions },
            });
        });
        const expectedActions = [
            { type: types.REQUEST_TRANSACTIONS },
            { type: types.RECEIVE_TRANSACTIONS, transactions: transactions }
        ]
        const store = mockStore({
            transactions: []
        })

        return store.dispatch(actions.getTransactions('1', '')).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
    it('creates RECEIVE_TRANSACTIONS when CR fetching transactions has been done', async () => {
        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.respondWith({
                status: 200,
                response: { transactions: transactions },
            });
        });
        const expectedActions = [
            { type: types.REQUEST_TRANSACTIONS },
            { type: types.RECEIVE_TRANSACTIONS, transactions: transactions }
        ]
        const store = mockStore({
            transactions: []
        })

        return store.dispatch(actions.getTransactions('1', 'CR')).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
    it('creates RECEIVE_TRANSACTIONS when fetching DR transactions has been done', async () => {
        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.respondWith({
                status: 200,
                response: { transactions: transactions },
            });
        });
        const expectedActions = [
            { type: types.REQUEST_TRANSACTIONS },
            { type: types.RECEIVE_TRANSACTIONS, transactions: transactions }
        ]
        const store = mockStore({
            transactions: []
        })

        return store.dispatch(actions.getTransactions('1', 'DR')).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
    it('creates RECEIVE_FAILURE when fetching transactions has been done', async () => {
        const errorResp = 'invalid data';

        moxios.wait(() => {
            const request = moxios.requests.mostRecent();
            request.reject(errorResp);
        });

        const expectedActions = [
            { type: types.REQUEST_TRANSACTIONS },
            { type: types.RECEIVE_FAILURE, error: errorResp }
        ]
        const store = mockStore({
            error: 'Error Occurred',
        })

        return store.dispatch(actions.getTransactions('1', '')).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })
})