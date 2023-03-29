import api from "./configs";

export const channelAPI = {
  getChannel: async id => {
    const { data } = await api.request({
      url: `/channel/${id}`,
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