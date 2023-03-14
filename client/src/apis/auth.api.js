import api from "./configs"

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
  }
}