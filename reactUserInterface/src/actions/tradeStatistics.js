import axios from "axios";

export const FETCH_ALL_TRADES_INITIAL = 'FETCH_ALL_TRADES_INITIAL';
export const FETCH_ALL_TRADES_REQUEST = 'FETCH_ALL_TRADES_REQUEST';
export const FETCH_ALL_TRADES_SUCCESS = 'FETCH_ALL_TRADES_SUCCESS';
export const FETCH_ALL_TRADES_FAILURE = 'FETCH_ALL_TRADES_FAILURE';
export const FETCH_MARKETSEARCH_FAILURE = 'FETCH_MARKETSEARCH_FAILURE';
export const FETCH_MARKETSEARCH_SUCCESS = 'FETCH_MARKETSEARCH_SUCCESS';
export const FETCH_MARKETSEARCH_REQUEST = 'FETCH_MARKETSEARCH_REQUEST';


function requestMarketSearch() {
  return {
    type: FETCH_MARKETSEARCH_REQUEST,
  };
}

function marketSearchError(message) {
  return {
    type: FETCH_MARKETSEARCH_FAILURE,
    message: message,
  };
}
//Get all related market search strings from the broker
function marketSearchSuccess(markets) {
  console.log("this is the trades inside marketseearch success", markets.length)
  return {
    type: FETCH_ALL_TRADES_SUCCESS,
    isFetching: false,
    markets,
  };
}

function requestFetchTrades() {
  return {
    type: FETCH_ALL_TRADES_REQUEST,
    isFetching: true,
  };
}

function fetchTradesSuccess(trades) {
  console.log("this is the trades inside fetch success", trades.longPositionPercentage)
  return {
    type: FETCH_ALL_TRADES_SUCCESS,
    isFetching: false,
    trades,
  };
}

function fetchTradesError(message) {
  return {
    type: FETCH_ALL_TRADES_FAILURE,
    isFetching: false,
    message: message,
  };
}



export function fetchTrades() {
  const config = {
    method: 'get',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  };
  return dispatch => {
    // We dispatch requestFetchTrades to kickoff the call to the API
    dispatch(requestFetchTrades());
    return fetch('http://localhost:9090/v1/igapi/clientsentiment/US500', config)
        .then(response => response.json().then(post => ({ post, response })))
        .then(({ post, response }) => {
          if (!response.ok) {
            // If there was a problem, we want to
            // dispatch the error condition
            dispatch(fetchTradesError(post.message));
            return Promise.reject(post);
          }
      //console.log(">>>>>>>>>>>>>>we are in service API for trades longPositionPercentage", response.data);
          dispatch(fetchTradesSuccess(post));
          return Promise.resolve(post);
        })
        .catch(err => console.error('Error: ', err));
  };
}

export const fetchSentiment = (searchTerm) => {
  const config = {
    method: 'get',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  };
  return dispatch => {
    // We dispatch requestFetchTrades to kickoff the call to the API
    dispatch(requestFetchTrades());
    return fetch('http://localhost:9090/v1/igapi/clientsentiment/'+searchTerm, config)
        .then(response => response.json().then(post => ({ post, response })))
        .then(({ post, response }) => {
          if (!response.ok) {
            // If there was a problem, we want to
            // dispatch the error condition
            dispatch(fetchTradesError(post.message));
            return Promise.reject(post);
          }
          //console.log(">>>>>>>>>>>>>>we are in service API for trades longPositionPercentage", response.data);
          dispatch(fetchTradesSuccess(post));
          return Promise.resolve(post);
        })
        .catch(err => console.error('Error: ', err));
  };
}

export function marketSearch(searchTerm) {
  const config = {
    method: 'get',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  };
  return dispatch => {
    // We dispatch requestFetchTrades to kickoff the call to the API
    dispatch(requestFetchTrades());
    return fetch('http://localhost:9090/v1/igapi/marketsearch/'+searchTerm, config)
        .then(response => response.json().then(post => ({ post, response })))
        .then(({ post, response }) => {
          if (!response.ok) {
            // If there was a problem, we want to
            // dispatch the error condition
            dispatch(fetchTradesError(post.message));
            return Promise.reject(post);
          }
          //console.log(">>>>>>>>>>>>>>we are in service API for trades longPositionPercentage", response.data);
          dispatch(fetchTradesSuccess(post));
          return Promise.resolve(post);
        })
        .catch(err => console.error('Error: ', err));
  };
}
