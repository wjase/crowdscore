import {LOAD_TEAMLIST_SUCCESS} from './constants';
import teamListApi from '../../api/teamListApi';

export function loadTeamsSuccess(teamList){
    return {type: LOAD_TEAMLIST_SUCCESS, teamList
    };
}

export const loadTeams = () => {

    return function(dispatch){

        return teamListApi.getAllTeams()
            .then(teams => {
                dispatch(loadTeamsSuccess(teams));
            })
        .catch(error => {throw error;});
    };

}
