import React from 'react';
import { connect } from 'react-redux';
import * as teamListActions from './actions';
import {bindActionCreators} from 'redux';
import List from 'grommet/components/List';
import ListItem from 'grommet/components/ListItem';


class TeamList extends React.Component {
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
    <div>Our worthy Teams
      <List>{this.props.teamList.map(team => (
        <ListItem key={team.ID}>Team: {team.Description}</ListItem>))}</List>
    </div>

  )};

}



const mapStateToProps = state => {
  return {
    teamList: state.teamList
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
)(TeamList);
