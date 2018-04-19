class RoundsAPI{
    static getAllTeams(){
    return fetch("/api/heats")
        .then(response=>{
            return response.json();
        })
        .catch(error=>{throw error});
    }
}

export default RoundsAPI
