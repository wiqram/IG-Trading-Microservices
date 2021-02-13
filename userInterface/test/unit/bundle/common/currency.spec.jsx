import React from 'react';
import { create } from 'react-test-renderer';
import CurrencyConverter from '../../../../src/bundle/common/component/currency';

describe('CurrencyConverter Snapshot Test', () => {

    it('renders correctly with all props given -> £', () => {
        const amount = 1212.22;
        const type = 'CR';
        const currency = '£'
        const wrapper = create(<CurrencyConverter amount={amount} type={type} currency={currency}/>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });
    it('renders correctly with all props given -> $', () => {
        const amount = 1212.22;
        const type = 'DR';
        const currency = '$'

        const wrapper = create(<CurrencyConverter amount={amount} type={type} currency={currency}/>)
            .toJSON();
        expect(wrapper).toMatchSnapshot();
    });
});