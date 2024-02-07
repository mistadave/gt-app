import { createSelector, createSlice } from "@reduxjs/toolkit";

const initialState = {
  user: JSON.parse(sessionStorage.getItem("user")),
  token: sessionStorage.getItem("token"),
};

export const loginSlice = createSlice({
  name: "login",
  initialState,
  reducers: {
    addUserToState: (state, action) => {
      state.user = action.payload.user;
      state.token = action.payload.token;
    },
    removeUserFromState: (state) => {
      state.user = {};
      state.token = "";
    },
  },
});

export const { addUserToState, removeUserFromState } = loginSlice.actions;
export const loginReducer = loginSlice.reducer;

const user = (state) => state.login.user;
const token = (state) => state.login.token;

export const checkAuthenticated = createSelector([user, token], (u, t) => {
  return u && t ? true : false;
});

export const getToken = createSelector([user, token], (u, t) => t);

export const getUser = createSelector([user, token], (u, t) => u);
