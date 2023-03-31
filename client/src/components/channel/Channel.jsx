import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { useLoaderData, useOutletContext, useParams } from "react-router-dom";
import { setChannel } from "../../store/slices/server.slice";
import ChannelChat from "./channel-chat/ChannelChat";
import ChannelNav from "./channel-nav/ChannelNav";
import ChannelUsersList from "./channel-users-list/ChannelUsersList";
import style from "./Channel.module.css";

const Channel = () => {
  const { id } = useParams()
  const dispatch = useDispatch()
  const channel = useLoaderData()
  const server = useOutletContext()

  useEffect(() => {
    localStorage.setItem(id, channel.id)
    dispatch(setChannel(channel))
  }, [channel, dispatch, id])

  return (
    <div className={style.container}>
      <ChannelNav name={channel.name} type={channel.type} topic={channel.topic} />
      <div className={style.content}>
        <ChannelChat key={"Chat" + channel.id} id={channel.id} messages={channel.messages} channelName={channel.name} users={server.users} />
        <ChannelUsersList key={"Users" + channel.id} channel={channel} users={server.users} roles={server.roles} />
      </div>
    </div >
  )
}

export default Channel