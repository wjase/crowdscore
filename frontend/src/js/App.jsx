//import styles
'use strict';
import 'grommet/scss/hpe/index';

//import { polyfill as promisePolyfill } from 'es6-promise';

import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import App from 'grommet/components/App';
import Sidebar from 'grommet/components/Sidebar';
import Box from 'grommet/components/Box';
import WorldMap from 'grommet/components/WorldMap';
import Header from 'grommet/components/Header';
import Menu from 'grommet/components/Menu';
import Anchor from 'grommet/components/Anchor';
import Footer from 'grommet/components/Footer';

import Meter from 'grommet/components/Meter';
import Title from 'grommet/components/Title';
import Value from 'grommet/components/Value';
import Split from 'grommet/components/Split';

import UserLogin from './components/Login'
import TeamList from './components/TeamList'
import Rounds from './components/Rounds'

import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import configureStore from './store/configure';
import {loadTeams} from './components/TeamList/actions';
import {loadRounds} from './components/Rounds/actions';

const store = configureStore({});

// initial rest fetch
store.dispatch(loadTeams());

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Switch from 'react-router-dom/Switch';

const About = () => (
  <div>
    <h2>About</h2>
  </div>
);

const Home = () => (
  <div>
    <h2>Home</h2>
    <Box pad='medium'>
      <Meter value={40} />
      <WorldMap value={40} />
    </Box>
  </div>

);

const Topics = ({ match }) => (
  <div>
    <h2>Topics</h2>
    <ul>
      <li>
        <Link to={`${match.url}/rendering`}>Rendering with React</Link>
      </li>
      <li>
        <Link to={`${match.url}/components`}>Components</Link>
      </li>
      <li>
        <Link to={`${match.url}/props-v-state`}>Props v. State</Link>
      </li>
    </ul>

    <Route path={`${match.url}/:topicId`} component={Topic} />
    <Route
      exact
      path={match.url}
      render={() => <h3>Please select a topic.</h3>}
    />
  </div>
);

class MainApp extends Component {
  constructor(props) {
    super(props);
    this.state = {showNav: true};
  }

  render() {
    return (
      <Provider store={store}>
      <App centered={false}>
        <Router>
          <Split
            fixed={false}
            flex='right'
            
          >
            <Sidebar pad='medium'>
              <Menu responsive={true}>
                <Link to='/app/'
                  className='active'>
                  Home
                </Link>
                <Link to='/app/teams'>
                  Teams
                </Link>
                <Link to='/app/rounds'>
                  Rounds
              </Link>
              </Menu>
            </Sidebar>
            <Box colorIndex='neutral-1'
              full='vertical'
              pad='medium'>
              <Header direction="row" size="medium"
                pad={{ horizontal: 'small' }}>
                {/* <button onClick={() => this.setState({showNav: true})}>Nav</button> */}
                <Title>CodeBlitz CrowdScore</Title>
              </Header>
              <Box
                full='vertical'
                pad='medium'>
                <Switch>
                  <Route exact path="/app/" component={Home} />
                  <Route path="/app/about" component={About} />
                  <Route path="/app/teams" component={TeamList} />
                  <Route path="/app/rounds" component={Rounds} />
                </Switch>
              </Box>
              <Footer fixed={true} primary={true} appCentered={false} direction="column"
                align="center" pad="small" colorIndex="grey-2"
              >
                <p>
                  CrowdScore - Letting the people speak.
          </p>
              </Footer>
            </Box>
          </Split>
        </Router>
      </App>
      </Provider>
    );
  }
};

let element = document.getElementById('app_content');
ReactDOM.render(React.createElement(MainApp), element);

document.body.classList.remove('loading');
