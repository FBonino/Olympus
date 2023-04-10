import api from "./configs";

export const authAPI = {
  login: async input => {
    const { data } = await api.request({
      url: "/auth/login",
      method: "POST",
      data: input
    })

    return data.user
  },
  signup: async input => {
    const { data } = await api.request({
      url: "/auth/signup",
      method: "POST",
      data: input
    })

    return data
  },
  logout: async () => {
    const { data } = await api.request({
      url: "/auth/logout",
      method: "POST"
    })

    return data
  },
  autoLogin: async () => {
    const { data } = await api.request({
      url: "/auth/auto-login",
      method: "GET"
    })

    return data.user
  }
}