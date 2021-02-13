import * as types from '../../../../../src/bundle/product/middleware/constants'
import productReducer from '../../../../../src/bundle/product/middleware/reducer'

describe('async products reducer', () => {
    const initialState = {
        isLoading: false,
        error: '',
        product: {},
        products: []
    };

    const products = [{
        productName: 'Saving A/C',
        currentBalance: '11.11',
        accountNumber: '111111',
        sortCode: '121212',
        availableBalance: '1212.12'
    }]
    it('returns correct state when action type is "DEFAULT"', () => {
        expect(productReducer(undefined, { type: 'DEFAULT' })).toEqual(initialState)
    })

    it('check initial state - INITIAL STATE', () => {

        const changeState = {
            isLoading: false,
            error: '',
            product: {},
            products: []
        };
        expect(productReducer(initialState, {})).toEqual(changeState);
    })

    it('check initial state - REQUEST_PRODUCTS ', () => {
        const startAction = {
            type: types.REQUEST_PRODUCTS
        };
        const changeState = {
            isLoading: true,
            error: '',
            product: {},
            products: []
        };
        expect(productReducer(initialState, startAction)).toEqual(changeState);
    })
    it('check initial state - RECEIVE_PRODUCTS', () => {
        const startAction = {
            type: types.RECEIVE_PRODUCTS,
            products: products
        };
        const changeState = {
            isLoading: false,
            error: '',
            product: {},
            products: products
        };
        expect(productReducer(initialState, startAction)).toEqual(changeState);
    })
    it('check initial state - RECEIVE_FAILURE', () => {
        const startAction = {
            type: types.RECEIVE_FAILURE,
            error: 'Error Occurred'
        };
        const changeState = {
            isLoading: false,
            error: 'Error Occurred',
            product: {},
            products: []
        };
        expect(productReducer(initialState, startAction)).toEqual(changeState);
    })
})