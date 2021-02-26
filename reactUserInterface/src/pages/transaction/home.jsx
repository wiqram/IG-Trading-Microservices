import React, { useEffect, useState } from 'react';
import { PropTypes } from 'prop-types';
import { useSelector, useDispatch } from 'react-redux';
import { Link, useLocation } from 'react-router-dom';
import { Col, ButtonPrimary } from '../../components/common/style/styles';
import CurrencyConverter from '../../components/common/component/currency';
import { getTransactions } from "../../actions/transaction"
import TransactionCard from './card';
import { CardRow, Card, FormRow, Label, ProductName, AccountNumber, TabButton } from './style';
import {Breadcrumb, BreadcrumbItem} from "reactstrap";

const TransactionHome = () => {
    const dispatch = useDispatch();
    const [tabs, setTabs] = useState(['ALL', 'IN', 'OUT']);
    const [activeTab, setActiveTab] = useState({ activeTab: 'ALL' })

    const useTransaction = useSelector(state => ({
        transactions: state.transactionState.transactions,
        isLoading: state.transactionState.isLoading,
        error: state.transactionState.error
    }));

    useEffect(() => {
        dispatch(getTransactions('SA001', 'ALL'));
    },[]);

    const getTransactionDetails = (e) => {
        e.preventDefault();
        let action = e.target.value;
        if (action === setTabs) return;
        setActiveTab(action);
        if (action === 'IN') {
            dispatch(getTransactions('SA001', 'CR'));
        } else if (action === 'OUT') {
            dispatch(getTransactions('SA001', 'DR'));
        } else {
            dispatch(getTransactions('SA001', 'ALL'));
        }

    };

    const { isLoading, transactions } = useTransaction;

    const { product } = useLocation();

        return (
            <div>
                <Breadcrumb>
                    <BreadcrumbItem>YOU ARE HERE</BreadcrumbItem>
                    <BreadcrumbItem active>Products > Transactions</BreadcrumbItem>
                </Breadcrumb>
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
                        tabs.map((tab) =>
                            <Col key={tab}>
                                <TabButton id={tab} value={tab} tabName={tab} isActive={setActiveTab}
                                           onClick={(e) => { getTransactionDetails(e) }}>{tab}</TabButton>
                            </Col>
                        )
                    }
                </CardRow>
                <TransactionCard isLoading={isLoading} transactions={transactions} />
            </div>
        );
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