class TeamListAPI{
    static getAllTeams(){
    return fetch("/api/teams")
        .then(response=>{
            return response.json();
        })
        .catch(error=>{throw error});
    }
}

export default TeamListAPI
