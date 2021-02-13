import * as types from '../../../../../src/bundle/transaction/middleware/constants'
import transactionReducer from '../../../../../src/bundle/transaction/middleware/reducer'

describe('async transactionReducer ', () => {
    const initialState = {
        isLoading: false,
        error: '',
        transactions: []
    };

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


    it('returns correct state when action type is "DEFAULT"', () => {
        expect(transactionReducer(undefined, { type: 'DEFAULT' })).toEqual(initialState)
    })

    it('check initial state - INITIAL STATE', () => {

        const changeState = {
            isLoading: false,
            error: '',
            transactions: []
        };
        expect(transactionReducer(initialState, {})).toEqual(changeState);
    })

    it('check initial state - REQUEST_TRANSACTIONS ', () => {
        const startAction = {
            type: types.REQUEST_TRANSACTIONS
        };
        const changeState = {
            isLoading: true,
            error: '',
            transactions: []
        };
        expect(transactionReducer(initialState, startAction)).toEqual(changeState);
    })
    it('check initial state - RECEIVE_TRANSACTIONS', () => {
        const startAction = {
            type: types.RECEIVE_TRANSACTIONS,
            transactions: transactions
        };
        const changeState = {
            isLoading: false,
            error: '',
            transactions: transactions
        };
        expect(transactionReducer(initialState, startAction)).toEqual(changeState);
    })
    it('check initial state - RECEIVE_FAILURE', () => {
        const startAction = {
            type: types.RECEIVE_FAILURE,
            error: 'Error Occurred'
        };
        const changeState = {
            isLoading: false,
            error: 'Error Occurred',
            transactions: []
        };
        expect(transactionReducer(initialState, startAction)).toEqual(changeState);
    })
})