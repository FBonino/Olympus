import api from "./configs";

export const channelAPI = {
  getChannel: async (id, limit) => {
    const { data } = await api.request({
      url: `/channel/${id}?limit=${limit}`,
      method: "GET"
    })

    return data.channel
  },
  newMessage: async (id, content) => {
    const { data } = await api.request({
      url: `/channel/${id}/messages`,
      method: "POST",
      data: { content }
    })

    return data.message
  }
}