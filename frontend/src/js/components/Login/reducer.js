import loginConstants from './constants';
import loginActions from './actions';

const reducer = (state = {}, action) => {
  switch (action.type) {
    case loginConstants.loaded:
    case loginConstants.submitted:
    default:
      return state;
  }
};

export default reducer;
