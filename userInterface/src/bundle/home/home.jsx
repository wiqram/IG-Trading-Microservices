import React from 'react';
import { Route, Switch, BrowserRouter as Router } from 'react-router-dom';
import ProductHome from '../product/home';
import TransactionHome from '../transaction/home';

export default class Home extends React.Component {
    render() {
        return (
            <Router>
                <Switch>
                    <Route path="/" exact component={ProductHome} />
                    <Route path="/product/:productId" exact component={TransactionHome} />
                </Switch>
            </Router>
        );
    }
}
