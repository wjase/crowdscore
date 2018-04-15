import React from 'react';
import { connect } from 'react-redux';

class Vote extends React.Component {
  static propTypes = {

  }
  render = () => (
    <div>
      Vote
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
)(Vote);
