import React from 'react';
import { create } from 'react-test-renderer';
import { Provider } from 'react-redux';
import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';
import { BrowserRouter } from 'react-router-dom';
import ProductHome from '../../../../src/bundle/product/home';

const middlewares = [thunk];
const mockStore = configureMockStore(middlewares);

describe('ProductHome Snapshot Test', () => {
    const products = [{
        id: '1',
        productId: "SA001",
        productName: 'Current A/C',
        currentBalance: 1121.11,
        accountNumber: '122111',
        sortCode: '122212',
        availableBalance: 12212.12

    },
    {
        id: '2',
        productId: "CA001",
        productName: 'Saving A/C',
        currentBalance: 11.11,
        accountNumber: '111111',
        sortCode: '121212',
        availableBalance: 1212.12

    }]
    let store;
    it('renders correctly with all props given error', () => {
        store = mockStore({
            productState: {
                isLoading: false,
                error: 'Test',
                products: []
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <ProductHome />
            </BrowserRouter>
        </Provider>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });
    it('renders correctly with all props given isLoading - false', () => {
        store = mockStore({
            productState: {
                isLoading: false,
                error: '',
                products: products
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <ProductHome />
            </BrowserRouter>
        </Provider>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });

    it('renders correctly with all props given isLoading - true', () => {
        store = mockStore({
            productState: {
                isLoading: true,
                error: '',
                products: []
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <ProductHome />
            </BrowserRouter>
        </Provider>)
        expect(wrapper).toMatchSnapshot();
    });
});