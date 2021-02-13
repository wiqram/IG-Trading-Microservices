import React from 'react';
import { create } from 'react-test-renderer';
import { BrowserRouter } from 'react-router-dom';
import ProductCard from '../../../../src/bundle/product/card';

describe('ProductCard Snapshot Test', () => {
    const product = {
        productName: 'Saving A/C',
        currentBalance: 11.11,
        accountNumber: '111111',
        sortCode: '121212',
        availableBalance: 1212.12

    }
    it('renders correctly with all props given', () => {
        const wrapper = create(<BrowserRouter><ProductCard product={product} /> </BrowserRouter>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });
});