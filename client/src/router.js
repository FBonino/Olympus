import { createBrowserRouter } from "react-router-dom";
import Me from "./components/me/Me";
import Auth from "./pages/auth/Auth";
import Home from "./pages/home/Home";
import Servers from "./pages/servers/Servers";
import Server from "./components/server/Server";
import Protected from "./components/protected/Protected";
import { serverAPI } from "./apis/server.api";
import ErrorLoading from "./components/error-loading/ErrorLoading";
import Channel from "./components/channel/Channel";
import Friends from "./components/me/friends/Friends";
import DirectMessage from "./components/me/direct-message/DirectMessage";
import { channelAPI } from "./apis/channels.api";
import { conversationAPI } from "./apis/conversation.api";
import { userAPI } from "./apis/user.api";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/auth",
    element: <Auth />
  },
  {
    path: "/",
    element: <Protected />,
    loader: async () => {
      const { user, servers } = await userAPI.getMyUser()
      return { user, servers }
    },
    children: [
      {
        path: "/channels",
        element: <Servers />,
        children: [
          {
            path: "/channels/@me",
            element: <Me />,
            loader: async () => {
              const conversations = await conversationAPI.me()
              return conversations
            },
            errorElement: <ErrorLoading />,
            children: [
              {
                path: "/channels/@me",
                element: <Friends />
              },
              {
                path: "/channels/@me/:conversation",
                element: <DirectMessage />,
                loader: async ({ params }) => {
                  const conversation = await conversationAPI.getConversation(params.conversation, 50)
                  return conversation
                }
              }
            ]
          },
          {
            path: "/channels/:server",
            element: <Server />,
            loader: async ({ params }) => {
              const server = await serverAPI.getServer(params.server)
              return server
            },
            errorElement: <ErrorLoading />,
            children: [
              {
                path: "/channels/:server/:channel",
                element: <Channel />,
                loader: async ({ params }) => {
                  const channel = await channelAPI.getChannel(params.channel, 50)
                  return channel
                },
              }
            ]
          }
        ],
      },
    ]
  },
])

export default router