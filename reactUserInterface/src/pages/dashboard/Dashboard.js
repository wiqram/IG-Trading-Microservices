import React, {useEffect, useState} from 'react';
import cx from 'classnames';
import PropTypes, {bool} from 'prop-types';
import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
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
  Table, InputGroup, Input, InputGroupAddon, NavItem
} from 'reactstrap';
import { mock } from './mock'

import Widget from '../../components/Widget';

import { fetchPosts } from '../../actions/posts';
import { fetchTrades, fetchSentiment, marketSearch } from '../../actions/tradeStatistics';
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
import Select from 'react-select'
import dashboardStyle from "../../assets/jss/material-dashboard-react/views/dashboardStyle";

const Dashboard = () => {
// call dispatch for actions
  const dispatch = useDispatch();
  const [selectOptions, setSelectOptions] = useState([]);
  const [id, setId] = useState("");
  const [name, setName] = useState("");
/*  const [classes, setClasses] = useState({});*/

// set state for the props
  const useDashboard = useSelector(state => ({
    isFetching: state.posts.isFetching,
    posts: state.posts.posts,
    classes: PropTypes.any.isRequired,
    trades: state.tradeStatistics.trades,
    markets: state.tradeStatistics.markets,
    marketSearchNames: state.tradeStatistics.marketSearchNames,
  }));



  // call reducer action now
  useEffect(() => {
    dispatch(fetchPosts());
    dispatch(fetchTrades());
    dispatch(marketSearch("US500"));
    }, []);


  // extract products list and others
  const { posts, trades, isFetching, markets, marketSearchNames, classes} = useDashboard;

    const formatDate = (str) => {
        return str.replace(/,.*$/,"");
    };

    /*const toggleDropdown = () => {
    this.setState(prevState => ({
      isDropdownOpened: !prevState.isDropdownOpened,
    }));
  };*/
  const handleChange = (e) => {
    console.log("this is inside handlechange  value--- ",e.value);
    console.log("this is inside handlechange  label--- ",e.label);
    setId(e.value);
    setName(e.label);
    dispatch(fetchSentiment(e.label));
  }

  const populateMarketSearchNames = () => {
    let arrayLength = markets.length;
    console.log("this is the market length --- ",markets.length);
    const options = [{
      "value" : "US500",
      "label" : "US500"
    }];
    for (let i = 0; i < arrayLength; i++) {
      console.log("this is the market EPIC --- ",markets[i].epic);
      let options1 = {
        "value" : markets[i].epic,
        "label" : markets[i].instrumentName
      };
      options.push(options1)
      //Do something
    }

    setSelectOptions(options);
  };

  const searchMarkets = (e) => {
    console.log("inside search markets - search term is - ", e.target.value);
    e.preventDefault();
    console.log("before calling dispath length of markets", markets.length);
    dispatch(marketSearch(e.target.value));
    console.log("after calling dispath length of markets", markets.length);
  };

  const searchSentiment = (e) => {
    console.log("inside search sentiment - search term is - ", e.target.value);
    e.preventDefault();
    dispatch(fetchSentiment(e.target.value));
  };


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
                  <NavItem>
                      <Select options={selectOptions} onChange={e => handleChange(e)}/>
                  </NavItem>
                      <p className={classes.cardCategory}>Active Trades</p>
                        <h3 className="mb-lg" color="warning">{trades.longPositionPercentage}
                          {trades.shortPositionPercentage}</h3>
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
                {posts && posts.map(post => (
                  <tr key={post.id}>
                    <td>{formatDate(new Date(post.updatedAt).toLocaleString())}</td>
                    <td>
                      <Link to="/app/posts">{post.title}</Link>
                    </td>
                  </tr>
                ))}
                {posts && !posts.length && (
                  mock.map(post => (
                    <tr key={post.id}>
                      <td>{post.updatedAt}</td>
                      <td>
                        <Link to="/app/posts">{post.title}</Link>
                      </td>
                    </tr>
                  ))
                )}
                {isFetching && (
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
      </div>
    );
  }

Dashboard.propTypes = {
  posts: PropTypes.any,
  trades: PropTypes.any,
  markets: PropTypes.any,
  marketSearchNames: PropTypes.any,
  isFetching: PropTypes.bool,
  classes: PropTypes.any.isRequired,
};


export default withStyles(dashboardStyle)(Dashboard);
