import React from 'react';
import { Link } from 'react-router-dom';
import { PropTypes } from 'prop-types';
import CurrencyConverter from '../common/component/currency';
import { Card, CardRow, ProductName, AccountNumber, Currency, BalanceLabel, NavItems } from './style';
import { Col, ButtonPrimary } from '../common/style/styles';

export default class ProductCard extends React.Component {
    render() {
        const { product } = this.props;
        return (

            <Card id={product.productId}>
                <Col col={9}>
                    <CardRow>
                        <ProductName id={product.productId+'_productName'}>{product.productName}</ProductName>
                        <AccountNumber id={product.productId+'_accountNumber'}>{product.accountNumber + '/' + product.sortCode}</AccountNumber>
                    </CardRow>
                    <CardRow>
                        <Currency size={36}>
                            <CurrencyConverter id={product.productId+'_currentBalanceValue'} amount={product.currentBalance} />
                        </Currency>
                        <BalanceLabel id={product.productId+'_currentBalanceLabel'}>Current Balance</BalanceLabel>
                    </CardRow>
                    <CardRow>
                        <Currency size={22}>
                            <CurrencyConverter id={product.productId+'_availableBalanceValue'} amount={product.availableBalance} />
                        </Currency>
                        <BalanceLabel id={product.productId+'_availableBalanceLabel'} >Available Balance</BalanceLabel>
                    </CardRow>
                </Col>
                <NavItems col={3}>
                    <Link id={product.productId+'_Link'}  to={{ pathname: `/product/${product.productId}`, product }}>
                        <ButtonPrimary>More Details</ButtonPrimary>
                    </Link>
                </NavItems>

            </Card>

        );
    }
}
ProductCard.propTypes = {
    product: PropTypes.object,
};