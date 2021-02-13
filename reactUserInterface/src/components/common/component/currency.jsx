import React, { Component } from 'react';
import { PropTypes } from 'prop-types';
import styled from 'styled-components';
import { colors } from '../style/base';

const TextStyle = styled.div`
    color:${props => (props.type === 'DR') ? colors.redColor : colors.greenColor};;
    `;

export default class CurrencyConverter extends Component {

    formatCurrency = (amountIn) => {
        const amount = (parseFloat(amountIn)).toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,');
        return amount;
    }
    render() {
        const { amount, type, currency, id } = this.props;
        return (

            <TextStyle id={id} type={type}>
                {!currency ? '£' + this.formatCurrency(amount) : '$' + this.formatCurrency(amount)}
            </TextStyle>

        );
    }
}
CurrencyConverter.propTypes = {
    amount: PropTypes.number,
    currency: PropTypes.string,
    type: PropTypes.string,
    id: PropTypes.string,
};