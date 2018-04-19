const reducer = (state = [], action) => {
  switch (action.type) {
    case "LOAD_TEAMLIST_SUCCESS":
      return action.teamList
    default:
      return state;
  }
};

export default reducer;
