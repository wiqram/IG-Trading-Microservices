import React from 'react';
import { create } from 'react-test-renderer';
import { Provider } from 'react-redux';
import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';
import { BrowserRouter } from 'react-router-dom';
import TransactionHome from '../../../../src/bundle/transaction/home';

const middlewares = [thunk];
const mockStore = configureMockStore(middlewares);

describe('TransactionHome Snapshot Test', () => {
    const transactions = [
        {
            id: 1,
            dateOn: "01-01-2019",
            transactionList: [
                {
                    id: 1,
                    productId: "SA001",
                    dateOn: "01-01-2019",
                    debitCredit: "CR",
                    transDescription: "ASDA Shopping",
                    details: "Trans at ADSA Livingston",
                    amount: 10010.2
                }
            ]
        },
        {
            id: 2,
            dateOn: "02-01-2019",
            transactionList: [
                {
                    id: 2,
                    productId: "SA001",
                    dateOn: "02-01-2019",
                    debitCredit: "DR",
                    transDescription: "MORRISON Shopping",
                    details: "Trans at MORRISON Livingston",
                    amount: 10020.2
                }
            ]
        }];
    const product = {
        productName: 'Saving A/C',
        currentBalance: 11.11,
        accountNumber: '111111',
        sortCode: '121212',
        availableBalance: 1212.12

    }
    const location = { product };
    const mockGetTransactionDetails = (e) => jest.fn();
    let store;
    it('renders correctly with all props given error', () => {
        store = mockStore({
            transactionState: {
                isLoading: false,
                error: 'Test',
                transactions: []
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <TransactionHome location={location} getTransactionDetails={mockGetTransactionDetails} />
            </BrowserRouter>
        </Provider>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });
    it('renders correctly with all props given isLoading - false', () => {
        store = mockStore({
            transactionState: {
                isLoading: false,
                error: '',
                transactions: transactions
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <TransactionHome location={location} getTransactionDetails={mockGetTransactionDetails} />
            </BrowserRouter>
        </Provider>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });

    it('renders correctly with all props given isLoading - true', () => {
        store = mockStore({
            transactionState: {
                isLoading: true,
                error: '',
                transactions: []
            }
        });
        const wrapper = create(<Provider store={store}>
            <BrowserRouter>
                <TransactionHome location={location} getTransactionDetails={mockGetTransactionDetails} />
            </BrowserRouter>
        </Provider>)
        expect(wrapper).toMatchSnapshot();
    });
});