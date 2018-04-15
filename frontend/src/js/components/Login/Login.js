import React from 'react';
import { connect } from 'react-redux';
import LoginForm from 'grommet/components/LoginForm';
import loginConstants from './constants'
import { bindActionCreators } from 'redux'
import loginActions from './actions'

class Login extends React.Component {
  static propTypes = {

  }

  constructor(props) {
    super(props);
  }
  render = () => {

    return (
      <LoginForm
        title="Login"
        errors={this.props.errors}
        onSubmit={this.props.submitted}
      />
    );

  }
}

function mapStateToProps(state) {

  return {

    errors: state.errors

  };

}

const mapDispatchToProps = dispatch => {
  let obj = bindActionCreators(loginActions, dispatch);
  return obj;
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login);
