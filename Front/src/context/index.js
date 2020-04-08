import React, { useReducer, createContext, useContext } from "react";

const userStateContext = createContext();
const userDispatchContext = createContext();

const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      return { ...state, user: action.user, token: action.token };
    case "LOGOUT":
      return { ...state, user: null, token: null };
    default:
      throw new Error(`Unknown action: ${action.type}`);
  }
};
/* eslint-disable */
export const UserProvider = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, {
    user: null,
    token: null
  });

  return (
    <userDispatchContext.Provider value={dispatch}>
      <userStateContext.Provider value={state}>
        {children}
      </userStateContext.Provider>
    </userDispatchContext.Provider>
  );
};
/* eslint-enable */

export const userContext = () => useContext(userStateContext);
export const userDispatch = () => useContext(userDispatchContext);
