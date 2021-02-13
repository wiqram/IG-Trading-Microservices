import React from "react";
import ReactDOM from "react-dom";
import { createBrowserHistory } from "history";
import {Router, Route, Switch, Redirect, BrowserRouter} from "react-router-dom";

// core components
import Admin from "layouts/Admin.jsx";
import Auth from "layouts/Auth.jsx";
import RTL from "layouts/RTL.jsx";

import "../resources/css/material-dashboard-react.css?v=1.6.0";
import store from "./store/index";
import {Provider} from "react-redux";


const hist = createBrowserHistory();

ReactDOM.render(
    <Router>
        <Provider store={store}>
            <Route path="/admin" component={Admin} />
            <Route path="/auth" component={Auth} />
            <Route path="/rtl" component={RTL} />
            <Redirect from="/" to="/admin/dashboard" />
        </Provider>
    </Router>,
  document.getElementById("root")
);
