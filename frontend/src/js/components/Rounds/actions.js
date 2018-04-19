import {LOAD_ROUNDS_SUCCESS} from './constants';
import roundsApi from '../../api/roundsApi';

export function loadRoundsSuccess(rounds){
    return {type: LOAD_ROUNDS_SUCCESS, rounds
    };
}

export const loadRounds = () => {

    return function(dispatch){

        return roundsApi.getAllRounds()
            .then(rounds => {
                dispatch(loadRoundsSuccess(rounds.map(x=>{
                    x.key =x.ID;
                    return x;
                })));
            })
        .catch(error => {throw error;});
    };

}
