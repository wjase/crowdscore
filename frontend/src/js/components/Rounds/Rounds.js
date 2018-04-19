import React from 'react';
import { connect } from 'react-redux';
import * as teamListActions from './actions';
import {bindActionCreators} from 'redux';
import List from 'grommet/components/List';
import ListItem from 'grommet/components/ListItem';


class Rounds extends React.Component {
  constructor(props, context) {
    super(props, context);
  }

  static propTypes = {

  };
  static defaultProps = {
    teamList: []
  }


  render = () => {
    return (
    <div>Rounds
      <List>{this.props.rounds.map(round => (
        <ListItem>{round.Description}</ListItem>))}</List>
    </div>

  )};

}



const mapStateToProps = state => {
  return {
    rounds: state.rounds
  };
};

const mapDispatchToProps = dispatch => (
   {
    actions: bindActionCreators(teamListActions, dispatch)
}
);
  


export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Rounds);
