//import styles
'use strict';
import 'grommet/scss/hpe/index';
import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import App from 'grommet/components/App';
import Sidebar from 'grommet/components/Sidebar';
import Box from 'grommet/components/Box';
import Header from 'grommet/components/Header';
import Menu from 'grommet/components/Menu';
import Footer from 'grommet/components/Footer';

import Meter from 'grommet/components/Meter';
import Button from 'grommet/components/Button';
import Title from 'grommet/components/Title';
import Split from 'grommet/components/Split';
import TeamList from './components/TeamList'
import Rounds from './components/Rounds'

import {Provider} from 'react-redux';
import configureStore from './store/configure';
import {loadTeams} from './components/TeamList/actions';
import {loadRounds} from './components/Rounds/actions';
import {BrowserRouter as Router, Link, Route} from "react-router-dom";
import Switch from 'react-router-dom/Switch';
import MenuIcon from 'grommet/components/icons/base/Menu';

//import { polyfill as promisePolyfill } from 'es6-promise';

const store = configureStore({});

// initial rest fetch
store.dispatch(loadTeams());
store.dispatch(loadRounds());

const About = () => (
    <div>
        <h2>About</h2>
    </div>
);

const Home = () => (
    <div>
        <h2>Home</h2>
        <Box pad='medium'>
            <Meter value={40}/>
            Boom. Codeblitz is great.
        </Box>
    </div>

);

class MainApp extends Component {
    constructor(props) {
        super(props);
        this.state = {showNavResponsive: false, isResponsive: false};
        this._onResponsive = this._onResponsive.bind(this);
        this._toggleSidebar = this._toggleSidebar.bind(this);
    }

    _onResponsive = function (isResponsiveNow) {
        console.log("Responsive: " + isResponsiveNow);
        this.setState(() => ({
            isResponsive: isResponsiveNow==='single'
        }));
    };

    _toggleSidebar = function(){
        this.setState({showNavResponsive: !this.state.showNavResponsive});
    };

    render() {

        const toggleSidebarButton = this.state.isResponsive ? (
            <Button onClick={this._toggleSidebar}><MenuIcon /></Button>
        ) : "";

        return (
            <Provider store={store}>
                <App centered={false}>
                    <Router>
                        <Split
                            fixed={false}
                            flex='right'
                            priority={this.state.showNavResponsive ? "left" : "right"}
                            onResponsive={this._onResponsive}
                        >
                            <Sidebar pad='medium'>
                                {toggleSidebarButton}
                                <Menu responsive={false}>
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
                                        pad={{horizontal: 'small'}}>
                                    {toggleSidebarButton}
                                    {/* <button onClick={() => this.setState({showNav: true})}>Nav</button> */}
                                    <Title>CodeBlitz CrowdScore</Title>
                                </Header>
                                <Box
                                    full='vertical'
                                    pad='medium'>
                                    <Switch>
                                        <Route exact path="/app/" component={Home}/>
                                        <Route path="/app/about" component={About}/>
                                        <Route path="/app/teams" component={TeamList}/>
                                        <Route path="/app/rounds" component={Rounds}/>
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
