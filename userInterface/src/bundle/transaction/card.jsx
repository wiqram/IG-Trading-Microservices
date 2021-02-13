import React from 'react';
import { PropTypes } from 'prop-types';
import { Col } from '../common/style/styles';
import CurrencyConverter from '../common/component/currency';
import Spinner from '../common/component/spinner';
import { TRow, OtherLabel, DateLabel } from './style';


class TransactionCard extends React.Component {

    render() {

        const { isLoading, transactions } = this.props;
        if (isLoading === true) {
            return <Spinner />
        }
        return (
            <div>
                {
                    transactions.map((transactionOn) =>
                        <div key={transactionOn.id}>
                            <TRow color={'white'}>
                                <Col><DateLabel>{transactionOn.dateOn}</DateLabel></Col>
                            </TRow>
                            {
                                transactionOn.transactionList.map((transaction) =>
                                    <TRow key={transaction.id}>
                                        <Col col={9}>
                                            <OtherLabel>{transaction.transDescription}</OtherLabel>
                                        </Col>
                                        <Col col={1}>
                                            <OtherLabel>{transaction.debitCredit}</OtherLabel>
                                        </Col>

                                        <Col col={3}>
                                            <OtherLabel>
                                                <CurrencyConverter type={transaction.debitCredit} amount={transaction.amount} />
                                            </OtherLabel>
                                        </Col>
                                    </TRow>
                                )
                            }
                        </div>
                    )
                }
            </div>
        );
    }
}
TransactionCard.propTypes = {
    transactions: PropTypes.array,
    isLoading: PropTypes.bool,
};
export default TransactionCard;