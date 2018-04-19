// The initial application state
import {combineReducers, createStore} from "redux";
import loginReducer from "../components/Login/reducer"
import roundsReducer from "../components/Rounds/reducer"
import teamListReducer from "../components/TeamList/reducer"



  export default combineReducers({
    rounds:roundsReducer,
    login:loginReducer,
    teamList:teamListReducer,
});
