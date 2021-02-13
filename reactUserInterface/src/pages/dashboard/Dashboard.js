import React, {Component} from 'react';
import cx from 'classnames';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import {
  Row,
  Col,
  Alert,
  Button,
  ButtonGroup,
  Breadcrumb,
  BreadcrumbItem,
  Progress,
  Badge,
  ListGroup,
  ButtonDropdown,
  DropdownMenu,
  DropdownToggle,
  DropdownItem,
  Table
} from 'reactstrap';
import { mock } from './mock'

import Widget from '../../components/Widget';

import { fetchPosts } from '../../actions/posts';
import { fetchTrades } from '../../actions/tradeStatistics';
import s from './Dashboard.module.scss';

import withStyles from "@material-ui/core/styles/withStyles";
import Icon from "@material-ui/core/Icon";
// @material-ui/icons
import Store from "@material-ui/icons/Store";
import Warning from "@material-ui/icons/Warning";
import DateRange from "@material-ui/icons/DateRange";
import LocalOffer from "@material-ui/icons/LocalOffer";
import Update from "@material-ui/icons/Update";
import ArrowUpward from "@material-ui/icons/ArrowUpward";
import AccessTime from "@material-ui/icons/AccessTime";
import Accessibility from "@material-ui/icons/Accessibility";
import BugReport from "@material-ui/icons/BugReport";
import Code from "@material-ui/icons/Code";
import Cloud from "@material-ui/icons/Cloud";
import GridContainer from "../../components/Grid/GridContainer";
import GridItem from "../../components/Grid/GridItem";
import Card from "../../components/Card/Card";
import CardHeader from "../../components/Card/CardHeader";
import CardIcon from "../../components/Card/CardIcon";
import CardFooter from "../../components/Card/CardFooter";
import Danger from "../../components/Typography/Danger";
import dashboardStyle from "../../assets/jss/material-dashboard-react/views/dashboardStyle";

class Dashboard extends Component {
  /* eslint-disable */
  static propTypes = {
    posts: PropTypes.any,
    trades: PropTypes.any,
    isFetching: PropTypes.bool,
    dispatch: PropTypes.func.isRequired,
    classes: PropTypes.object.isRequired,
  };
  /* eslint-enable */

  static defaultProps = {
    posts: [],
    longPositionPercentage: "",
    shortPositionPercentage: "",
    isFetching: false,
    classes: [],
  };

  state = {
    isDropdownOpened: false
  };

  componentDidMount() {
    if(process.env.NODE_ENV === "development") {
      this.props.dispatch(fetchPosts());
      this.props.dispatch(fetchTrades());
    }
  }

  formatDate = (str) => {
    return str.replace(/,.*$/,"");
  }

  toggleDropdown = () => {
    this.setState(prevState => ({
      isDropdownOpened: !prevState.isDropdownOpened,
    }));
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={s.root}>
        <Breadcrumb>
          <BreadcrumbItem>YOU ARE HERE</BreadcrumbItem>
          <BreadcrumbItem active>Dashboard</BreadcrumbItem>
        </Breadcrumb>
        <h1 className="mb-lg">Dashboard</h1>

          <GridContainer>
            <GridItem xs={12} sm={6} md={3}>
              <Card>
                <CardHeader color="warning" stats icon>
                  <CardIcon color="warning">
                    <Icon>add_circle</Icon>
                  </CardIcon>
                  <p className={classes.cardCategory}>Active Trades</p>
                  <h3 className={classes.cardTitle}>
                    <Link to="/app/posts">{this.props.trades.longPositionPercentage}</Link>
                  </h3>
                </CardHeader>
                <CardFooter stats>
                  <div className={classes.stats}>
                    <Danger>
                      <Warning />
                    </Danger>
                    <a href="#pablo" onClick={e => e.preventDefault()}>
                      Get more space
                    </a>
                  </div>
                </CardFooter>
              </Card>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
              <Card>
                <CardHeader color="success" stats icon>
                  <CardIcon color="success">
                    <Store />
                  </CardIcon>
                  <p className={classes.cardCategory}>P&L</p>
                  <h3 className={classes.cardTitle}>$34,245</h3>
                </CardHeader>
                <CardFooter stats>
                  <div className={classes.stats}>
                    <DateRange />
                   Till Date
                  </div>
                </CardFooter>
              </Card>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
              <Card>
                <CardHeader color="danger" stats icon>
                  <CardIcon color="danger">
                    <Icon>info_outline</Icon>
                  </CardIcon>
                  <p className={classes.cardCategory}>Total Trades Executed</p>
                  <h3 className={classes.cardTitle}>75</h3>
                </CardHeader>
                <CardFooter stats>
                  <div className={classes.stats}>
                    <LocalOffer />
                    Tracked from Github
                  </div>
                </CardFooter>
              </Card>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
              <Card>
                <CardHeader color="info" stats icon>
                  <CardIcon color="info">
                    <Accessibility />
                  </CardIcon>
                  <p className={classes.cardCategory}>Followers</p>
                  <h3 className={classes.cardTitle}>+245</h3>
                </CardHeader>
                <CardFooter stats>
                  <div className={classes.stats}>
                    <Update />
                    Just Updated
                  </div>
                </CardFooter>
              </Card>
            </GridItem>
          </GridContainer>

{/*        <Row>
          <Col sm={12} md={6}>
            <Widget
              title={
                <div>
                  <div className="pull-right mt-n-xs">
                    <input
                      type="search"
                      placeholder="Search..."
                      className="form-control input-sm"
                    />
                  </div>
                  <h5 className="mt-0 mb-3">
                    <i className="fa fa-user mr-xs opacity-70" />{' '}
                    Users
                  </h5>
                </div>
              }
            >
              <Table responsive borderless className={cx('mb-0', s.usersTable)}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Username</th>
                    <th>Email</th>
                    <th>Status</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>1</td>
                    <td>Alice</td>
                    <td>alice@email.com</td>
                    <td>
                      <span className="py-0 px-1 bg-success rounded text-white">active</span>
                    </td>
                  </tr>
                  <tr>
                    <td>2</td>
                    <td>Bob</td>
                    <td>bob@email.com</td>
                    <td>
                      <span className="py-0 px-1 bg-warning rounded text-white">delayed</span>
                    </td>
                  </tr>
                  <tr>
                    <td>3</td>
                    <td>Duck</td>
                    <td>duck@email.com</td>
                    <td>
                      <span className="py-0 px-1 bg-success rounded text-white">active</span>
                    </td>
                  </tr>
                  <tr>
                    <td>4</td>
                    <td>Shepherd</td>
                    <td>shepherd@email.com</td>
                    <td>
                      <span className="py-0 px-1 bg-danger rounded text-white">removed</span>
                    </td>
                  </tr>
                </tbody>
              </Table>
            </Widget>
          </Col>
          <Col sm={12} md={6}>
            <Widget title="Alerts">
              <Alert
                className="alert-sm"
                color="warning"
              >
                <span className="fw-semi-bold">Warning:</span> Best check yo
                self, you&#39;re not looking too good.
              </Alert>
              <Alert
                className="alert-sm"
                color="success"
              >
                <span className="fw-semi-bold">Success:</span> You successfully
                read this important alert message.
              </Alert>
              <Alert
                className="alert-sm"
                color="info"
              >
                <span className="fw-semi-bold">Info:</span> This alert needs
                your attention, but it&#39;s not super important.
              </Alert>
              <Alert
                className="alert-sm clearfix"
                color="danger"
              >
                <span className="fw-semi-bold">Danger:</span> Change this and
                that and try again.
                <span className="pull-right mr-sm">
                  <Button color="danger" size="sm">
                    Take this action
                  </Button>
                  <span className="px-2"> or </span>
                  <Button color="default" size="sm">Cancel</Button>
                </span>
              </Alert>
            </Widget>
          </Col>
        </Row>*/}
        <Row>
          <Col sm={6}>
            <Widget
              title={
                <div>
                  <div className="pull-right mt-n-xs">
                    <Link to="/app/main" className="td-underline fs-sm">Options</Link>
                  </div>
                  <h5 className="mt-0 mb-0">
                    Recent posts{' '}
                    <Badge color="success" className="ml-xs">
                      5
                    </Badge>
                  </h5>
                  <p className="fs-sm mb-0 text-muted">
                    posts, that have been published recently
                  </p>
                </div>
              }
            >
              <table className="table table-sm table-no-border mb-0">
                <tbody>
                {this.props.posts &&
                this.props.posts.map(post => (
                  <tr key={post.id}>
                    <td>{this.formatDate(new Date(post.updatedAt).toLocaleString())}</td>
                    <td>
                      <Link to="/app/posts">{post.title}</Link>
                    </td>
                  </tr>
                ))}
                {this.props.posts &&
                !this.props.posts.length && (
                  mock.map(post => (
                    <tr key={post.id}>
                      <td>{post.updatedAt}</td>
                      <td>
                        <Link to="/app/posts">{post.title}</Link>
                      </td>
                    </tr>
                  ))
                )}
                {this.props.isFetching && (
                  <tr>
                    <td colSpan="100">Loading...</td>
                  </tr>
                )}
                </tbody>
              </table>
              <div className="d-flex justify-content-end">
                <Link to="/app/posts" className="btn btn-default">
                  View all Posts <Badge className="ml-xs" color="danger">13</Badge>
                </Link>
              </div>
            </Widget>
          </Col>
          <Col sm={6}>
            <ListGroup>
              <Link to="/app" className="list-group-item">
                <i className="fa fa-phone mr-xs text-secondary" />{' '}
                Incoming calls <Badge className="ml-xs" color="danger">3</Badge>
              </Link>
              <Link to="/app" className="list-group-item">
                <i className="fa fa-bell-o mr-xs text-secondary" />{' '}
                Notifications <Badge className="ml-xs" color="warning">6</Badge>
              </Link>
              <Link to="/app" className="list-group-item">
                <i className="fa fa-comment-o mr-xs text-secondary" />{' '}
                Messages <Badge className="ml-xs" color="success">18</Badge>
              </Link>
              <Link to="/app" className="list-group-item">
                <i className="fa fa-eye mr-xs text-secondary" />{' '}
                Visits total
              </Link>
              <Link to="/app" className="list-group-item">
                <i className="fa fa-cloud mr-xs text-secondary" /> Inbox{' '}
              </Link>
            </ListGroup>
          </Col>
        </Row>
        <Widget className="mt-lg" title="Some standard reactstrap components">
          <Row>
            <Col sm={6}>
              <div className="mt">
                <Button size="sm" color="default" className="mr-sm mb-xs">
                  Default
                </Button>
                <Button size="sm" color="success" className="mr-sm mb-xs">
                  Success
                </Button>
                <Button size="sm" color="info" className="mr-sm mb-xs">
                  Info
                </Button>
                <Button size="sm" color="warning" className="mr-sm mb-xs">
                  Warning
                </Button>
                <Button size="sm" color="danger" className="mb-xs">
                  Danger
                </Button>
              </div>
              <ButtonGroup className="mb">
                <Button color="default">1</Button>
                <Button color="default">2</Button>
                <ButtonDropdown isOpen={this.state.isDropdownOpened} toggle={this.toggleDropdown}>
                  <DropdownToggle color="default" caret>
                    Dropdown
                  </DropdownToggle>
                  <DropdownMenu>
                    <DropdownItem>1</DropdownItem>
                    <DropdownItem>2</DropdownItem>
                  </DropdownMenu>
                </ButtonDropdown>
              </ButtonGroup>
              <p className="mb-0">
                For more components please checkout{' '}
                <a
                  href="https://reactstrap.github.io/"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  reactstrap documentation
                </a>
              </p>
            </Col>
            <Col sm={6}>
              <Progress className="progress-sm" color="success" value={40} />
              <Progress className="progress-sm" color="info" value={20} />
              <Progress className="progress-sm" color="warning" value={60} />
              <Progress className="progress-sm" color="danger" value={80} />
            </Col>
          </Row>
        </Widget>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    isFetching: state.posts.isFetching,
    posts: state.posts.posts,
    classes: PropTypes.object.isRequired,
    trades: state.tradeStatistics.trades,
  };
}

export default connect(mapStateToProps)(withStyles(dashboardStyle)(Dashboard));
