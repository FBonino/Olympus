import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./slices/user.slice";
import serverReducer from "./slices/server.slice";

const store = configureStore({
  reducer: {
    user: userReducer,
    server: serverReducer
  }
});

export default store
