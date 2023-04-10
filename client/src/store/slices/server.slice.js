import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { serverAPI } from "../../apis/server.api";
import { logout } from "./user.slice";

export const createServer = createAsyncThunk("server/create", async input => {
  const server = await serverAPI.createServer(input)
  return server
})

const initialState = {
  servers: [],
  server: null,
  channel: null
}

const serverSlice = createSlice({
  name: "server",
  initialState,
  reducers: {
    setServer: (state, { payload }) => {
      state.server = payload
    },
    setServers: (state, { payload }) => {
      state.servers = payload
    },
    setChannel: (state, { payload }) => {
      state.channel = payload
    },
    clearSelection: state => {
      return { ...initialState, servers: state.servers }
    }
  },
  extraReducers: builder => {
    builder
      .addCase(logout.fulfilled, (state, { payload }) => {
        return initialState
      })
      .addCase(createServer.fulfilled, (state, { payload }) => {
        state.servers.push(payload)
      })
  }
})

const { reducer } = serverSlice

export const { setServer, setServers, setChannel, clearSelection } = serverSlice.actions

export default reducer