import { authAPI } from "../../apis/auth.api";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export const logout = createAsyncThunk("user/logout", async () => {
  const data = await authAPI.logout()
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
  friends: []
}

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (state, { payload }) => {
      return payload
    }
  },
  extraReducers: builder => {
    builder
      .addCase(logout.fulfilled, (state, { payload }) => {
        return initialState
      })
  }
})

const { reducer } = userSlice

export const { setUser } = userSlice.actions

export default reducer