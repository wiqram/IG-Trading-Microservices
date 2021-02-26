/**
 * Flatlogic Dashboards (https://flatlogic.com/admin-dashboards)
 *
 * Copyright © 2015-present Flatlogic, LLC. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE.txt file in the root directory of this source tree.
 */

import React from 'react';
import {Breadcrumb, BreadcrumbItem, Row} from "reactstrap";

/*const __html = require('./app/src/main/web/index.html');
const template = { __html: __html };*/

class Igtrade extends React.Component {
    render() {
        return (
            <div>
                <Breadcrumb>
                    <BreadcrumbItem>YOU ARE HERE</BreadcrumbItem>
                    <BreadcrumbItem active>IG Trade</BreadcrumbItem>
                </Breadcrumb>
                <Row>
                    <div className="navbar navbar-inverse">
                        <div>
                            <div className="container">
                                <button type="button" className="btn btn-navbar" data-toggle="collapse"
                                        data-target=".nav-collapse">
                                    <span className="icon-bar"></span>
                                    <span className="icon-bar"></span>
                                    <span className="icon-bar"></span>
                                </button>
                                <div>
                                    <form className="navbar-form pull-left">
                                        Sample application
                                    </form>
                                    <form className="navbar-form pull-right">
                                        <input title="IG client login name" className="span2" type="text" id="username"
                                               placeholder="Username"/>
                                        <input title="IG client password" className="span2" type="password"
                                               id="password"
                                               placeholder="Password"/>
                                        <input title="Application key" className="span2" type="text" id="apikey"
                                               placeholder="API key"/>
                                        <input title="Login using IG client credentials" type="button" id="loginButton"
                                               className="button-green" value="Login"/>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </Row>
                <Row>
                    <div>
                        <h4>Open Positions</h4>
                        <input type="button" id="positionsButton" value="Retrieve Positions"/>

                        <table id="positions_list" className="table table-hover">
                            <thead>
                            <tr>
                                <th width="30px">Status</th>
                                <th>Market</th>
                                <th>Expiry</th>
                                <th>Size</th>
                                <th>Bid</th>
                                <th>Offer</th>
                            </tr>
                            </thead>
                            <tbody>
                            </tbody>
                        </table>

                        <h4>Search Markets</h4>
                        <input type="text" id="searchTerm"/>
                        <input type="button" id="searchButton" value="Search"/>

                        <table id="search_results_list" className="table table-hover">
                            <thead>
                            <tr>
                                <th width="30px">Status</th>
                                <th>Market</th>
                                <th>Expiry</th>
                                <th>Bid</th>
                                <th>Offer</th>
                            </tr>
                            </thead>
                            <tbody>
                            </tbody>
                        </table>

                        <script src="./app/src/main/javascript/jquery.js"></script>
                        <script src="./app/src/main/javascript/jquery-ui-1.10.3.custom.js"></script>
                        <script src="./app/src/main/require.js"></script>
                        <script src="./app/src/main/javascript/lightstreamer.js"></script>
                        <script src="./lib/src/main/javascript/ig-public-api.js"></script>
                        <script src="./app/src/main/javascript/pidder.js"></script>
                        <script src="./app/src/main/javascript/beautifier.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-transition.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-alert.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-modal.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-dropdown.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-scrollspy.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-tab.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-tooltip.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-popover.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-button.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-collapse.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-carousel.js"></script>
                        <script src="./app/src/main/javascript/bootstrap-typeahead.js"></script>
                        <script type="text/javascript" src="/sites/all/themes/IG/js/AppMeasurement.js"></script>
                    </div>
                </Row>
            </div>


        );
    }
}

export default Igtrade;
