import api from "./configs";

export const userAPI = {
  uploadAvatar: async image => {
    const { data } = await api.request({
      url: "/user/upload",
      method: "POST",
      data: image
    })

    return data.avatar
  },
}