import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { authAPI } from "../../apis/auth.api";

export const login = createAsyncThunk("user/login", async input => {
  const data = await authAPI.login(input)
  return data
})

export const logout = createAsyncThunk("user/logout", async () => {
  const data = await authAPI.logout()
  return data
})

export const autoLogin = createAsyncThunk("user/autoLogin", async input => {
  const data = await authAPI.autoLogin()
  return data
})

const initialState = {
  id: "",
  username: "",
  email: "",
  avatar: "",
  status: "",
  signedin: false,
  customStatus: "",
  servers: []
}

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(login.fulfilled, (state, { payload }) => {
        return { ...payload, signedin: true }
      })
      .addCase(login.rejected, () => {
        throw Error("Username/Email or Password is wrong!")
      })
      .addCase(logout.fulfilled, (state, { payload }) => {
        return initialState
      })
      .addCase(autoLogin.fulfilled, (state, { payload }) => {
        return { ...payload, signedin: true }
      })
  }
})

const { reducer } = userSlice

export default reducer