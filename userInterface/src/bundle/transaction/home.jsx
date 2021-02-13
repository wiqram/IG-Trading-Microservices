import React from 'react';
import { PropTypes } from 'prop-types';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Col, ButtonPrimary } from '../common/style/styles';
import CurrencyConverter from '../common/component/currency';
import { getTransactions } from "./middleware/actions"
import TransactionCard from './card';
import { CardRow, Card, FormRow, Label, ProductName, AccountNumber, TabButton } from './style';


@connect((state) => {
    return {
        transactions: state.transactionState.transactions,
        isLoading: state.transactionState.isLoading,
        error: state.transactionState.error
    };
})
class TransactionHome extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            activeTab: 'ALL',
            tabs: [
                'ALL',
                'IN',
                'OUT'
            ]
        }
    }
    componentDidMount() {
        console.log(this.props.location.product);
        
        this.props.dispatch(getTransactions('SA001', 'ALL'));
    }

    getTransactionDetails = (e) => {
        e.preventDefault();
        let action = e.target.value;
        if (action === this.state.activeTab) return;
        this.setState({ activeTab: action });

        if (action === 'IN') {
            this.props.dispatch(getTransactions('SA001', 'CR'));
        } else if (action === 'OUT') {
            this.props.dispatch(getTransactions('SA001', 'DR'));
        } else {
            this.props.dispatch(getTransactions('SA001', 'ALL'));
        }

    };

    render() {
        const product = this.props.location.product;
        const { isLoading, transactions } = this.props;


        return (
            <div>
                <CardRow>
                    <Col alignItems={'flex-start'}>
                        <Link to={`/`}><ButtonPrimary> &lt; Back to Accounts</ButtonPrimary> </Link>
                    </Col>
                </CardRow>
                <Card>
                    <FormRow>
                        <Col>
                            <Label id="accountNameLabel">Account Name</Label>
                        </Col>
                        <Col>
                            <ProductName id="accountNameValue">{product.productName}</ProductName>
                        </Col>
                    </FormRow>
                    <FormRow>
                        <Col>
                            <Label id="accountNumberLabel">Account Number</Label>
                        </Col>
                        <Col>
                            <AccountNumber id="accountNumberValue">{product.accountNumber}</AccountNumber>
                        </Col>
                    </FormRow>
                    <FormRow>
                        <Col>
                            <Label id="sortCodeLabel">Sort Code</Label>
                        </Col>
                        <Col>
                            <AccountNumber id="sortCodeValue">{product.sortCode}</AccountNumber>
                        </Col>
                    </FormRow>
                    <FormRow>
                        <Col>
                            <Label id="currentBalanceLabel">Current Balance</Label>
                        </Col>
                        <Col>
                            <CurrencyConverter id="currentBalanceValue" amount={product.currentBalance} />
                        </Col>
                    </FormRow>
                    <FormRow>
                        <Col>
                            <Label id="availableBalanceLabel">Available Balance</Label>
                        </Col>
                        <Col>
                            <CurrencyConverter id="availableBalanceValue" amount={product.availableBalance} />
                        </Col>
                    </FormRow>
                </Card>

                <CardRow>
                    <Label id="infoText">Click on tab to see your latest transaction</Label>
                </CardRow>

                <CardRow>
                    {
                        this.state.tabs.map((tab) =>
                            <Col key={tab}>
                                <TabButton id={tab} value={tab} tabName={tab} isActive={this.state.activeTab}
                                    onClick={(e) => { this.getTransactionDetails(e) }}>{tab}</TabButton>
                            </Col>
                        )
                    }
                </CardRow>
                <TransactionCard isLoading={isLoading} transactions={transactions} />
            </div>
        );
    }
}
TransactionHome.propTypes = {
    product: PropTypes.object,
    location: PropTypes.object,
    transactions: PropTypes.object,
    isLoading: PropTypes.bool,
    error: PropTypes.string,
    dispatch: PropTypes.func,
};
export default TransactionHome;