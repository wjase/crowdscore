 import loginConstants from './constants';
 import loginService from '../../services/login'

const loginLoaded = () => {
    return {type:loginConstants.loaded}
};

const loginSucceeded = (user) => {
    return {type:loginConstants.succeeded}
};

const submitted = ({username,password}) => {
    return dispatch => 
    {
        loginService.login(username,password)
            .then(
                user =>{
                    dispatch(loginSucceeded(user));
                },
                error =>{
                    dispatch(failure(error))
                    dispatch(loginActions.error(error))
                }
                
            );
    }
};

export default {
    loaded:loginLoaded,
    submitted:submitted
}
