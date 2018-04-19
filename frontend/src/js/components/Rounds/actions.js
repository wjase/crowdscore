import {LOAD_TEAMLIST_SUCCESS} from './constants';
import roundsApi from '../../api/roundsApi';

export function loadRoundsSuccess(rounds){
    return {type: LOAD_ROUNDS_SUCCESS, rounds
    };
}

export const loadRounds = () => {

    return function(dispatch){

        return roundsApi.getAllRounds()
            .then(rounds => {
                dispatch(loadRoundsSuccess(rounds));
            })
        .catch(error => {throw error;});
    };

}
