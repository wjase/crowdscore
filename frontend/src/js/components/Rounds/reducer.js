const reducer = (state = [], action) => {
  switch (action.type) {
    case "LOAD_ROUNDS_SUCCESS":
      return action.rounds
    default:
      return state;
  }
};

export default reducer;
