import api from "./configs";

export const serverAPI = {
  createServer: async input => {
    const { data } = await api.request({
      url: "/server",
      method: "POST",
      data: input
    })

    return data.server
  },
  getServer: async id => {
    const { data } = await api.request({
      url: `/server/${id}`,
      method: "GET"
    })

    return data.server
  },
  getChannel: async (id, channel) => {
    const { data } = await api.request({
      url: `/server/${id}/${channel}`,
      method: "GET"
    })

    return data.channel
  }
}