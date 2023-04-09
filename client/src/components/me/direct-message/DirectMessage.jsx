import React from "react";
import style from "./DirectMessage.module.css";
import DirectMessageNav from "../direct-message-nav/DirectMessageNav";
import { useLoaderData } from "react-router-dom";
import ChannelChat from "../../channel/channel-chat/ChannelChat";
import { conversationAPI } from "../../../apis/conversation.api";

const DirectMessage = () => {
  const conversation = useLoaderData()
  const { username, status } = conversation.users[0]

  return (
    <div className={style.container}>
      <DirectMessageNav status={status} username={username} />
      <div className={style.content}>
        <ChannelChat id={conversation.id} messages={conversation.messages} users={[...conversation.users, conversation.me]} key={conversation.id} createMessage={conversationAPI.newMessage} />
      </div>
    </div>
  )
}

export default DirectMessage