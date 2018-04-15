import React from 'react';
import { connect } from 'react-redux';

class Leaderboard extends React.Component {
  static propTypes = {

  }
  render = () => (
    <div>
      Leaderboard
    </div>
  )
}

const mapStateToProps = state => ({

});

const mapDispatchToProps = dispatch => ({

});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Leaderboard);
