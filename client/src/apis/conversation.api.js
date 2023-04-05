import api from "./configs";

export const conversationAPI = {
  me: async () => {
    const { data } = await api.request({
      url: "/conversation/me",
      method: "GET"
    })

    return data.conversations ?? []
  },
  create: async users => {
    const { data } = await api.request({
      url: "/conversation",
      method: "POST",
      data: { users }
    })

    return data.conversation
  },
  getConversation: async (id, limit) => {
    const { data } = await api.request({
      url: `/conversation/${id}?limit=${limit}`,
      method: "GET"
    })

    return data.conversation
  },
  newMessage: async (id, content) => {
    const { data } = await api.request({
      url: `/conversation/${id}/messages`,
      method: "POST",
      data: { content }
    })

    return data.message
  }
}