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
    children: [
      {
        path: "/channels",
        element: <Servers />,
        children: [
          {
            path: "/channels/@me",
            element: <Me />,
            errorElement: <ErrorLoading />,
            children: [
              {
                path: "/channels/@me",
                element: <Friends />
              },
              {
                path: "/channels/@me/:id",
                element: <DirectMessage />
              }
            ]
          },
          {
            path: "/channels/:id",
            element: <Server />,
            loader: async ({ request, params }) => {
              const server = await serverAPI.getServer(params.id)
              return server
            },
            errorElement: <ErrorLoading />,
            children: [
              {
                path: "/channels/:id/:channel",
                element: <Channel />,
                loader: async ({ request, params }) => {
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