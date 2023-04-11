import React from "react";
import style from "./DirectMessage.module.css";
import DirectMessageNav from "../direct-message-nav/DirectMessageNav";
import { useLoaderData } from "react-router-dom";
import Chat from "../../chat/Chat";
import { conversationAPI } from "../../../apis/conversation.api";

const DirectMessage = () => {
  const conversation = useLoaderData()
  const { username, status } = conversation.users.length ? conversation.users[0] : conversation.me

  return (
    <div className={style.container}>
      <DirectMessageNav status={status} username={username} />
      <div className={style.content}>
        <Chat id={conversation.id} messages={conversation.messages} users={[...conversation.users, conversation.me]} key={conversation.id} createMessage={conversationAPI.newMessage} />
      </div>
    </div>
  )
}

export default DirectMessage