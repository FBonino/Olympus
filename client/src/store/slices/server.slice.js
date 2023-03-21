import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { serverAPI } from "../../apis/server.api";
import { autoLogin, login, logout } from "./user.slice";

export const createServer = createAsyncThunk("server/create", async input => {
  const server = await serverAPI.createServer(input)
  return server
})

const initialState = {
  servers: []
}

const serverSlice = createSlice({
  name: "server",
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(login.fulfilled, (state, { payload }) => {
        return { servers: payload.servers }
      })
      .addCase(logout.fulfilled, (state, { payload }) => {
        return initialState
      })
      .addCase(autoLogin.fulfilled, (state, { payload }) => {
        return { servers: payload.servers }
      })
      .addCase(createServer.fulfilled, (state, { payload }) => {
        state.servers.push(payload)
      })
  }
})

const { reducer } = serverSlice

export default reducer