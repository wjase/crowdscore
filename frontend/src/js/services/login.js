function login(username,password){

    return {
        then : (happy,sad)=>{
            console.log('calling login callback');
            sad({username:username, password:password});
        }
    }

}

export default {
    login:login
};