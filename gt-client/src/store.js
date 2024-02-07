import { configureStore } from "@reduxjs/toolkit";
import { loginReducer } from "./redux/loginSlice";
import { transactionReducer } from "./redux/transactionSlice";

const store = configureStore({
  reducer: { login: loginReducer, transaction: transactionReducer },
});

export default store;
