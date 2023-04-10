import api from "./configs";

export const userAPI = {
  getMyUser: async () => {
    const { data } = await api.request({
      url: "/user/me",
      method: "GET"
    })

    localStorage.setItem("account", JSON.stringify(data.user))

    return data
  },
  uploadAvatar: async image => {
    const { data } = await api.request({
      url: "/user/upload",
      method: "POST",
      data: image
    })

    return data.avatar
  },
}