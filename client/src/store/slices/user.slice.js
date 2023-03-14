import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { authAPI } from "../../apis/auth.api";

export const login = createAsyncThunk("user/login", async input => {
  const user = await authAPI.login(input)
  return user
})

const userSlice = createSlice({
  name: "user",
  initialState: {
    id: "",
    username: "",
    email: "",
    avatar: "",
    status: "",
    signedin: false,
    customStatus: "",
  },
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(login.fulfilled, (state, { payload }) => {
        return payload
      })
      .addCase(login.rejected, () => {
        throw Error("Username/Email or Password is wrong!")
      })
  }
})

const { reducer } = userSlice

export default reducer