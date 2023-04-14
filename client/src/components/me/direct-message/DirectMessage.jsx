import React from "react";
import style from "./DirectMessage.module.css";
import DirectMessageNav from "../direct-message-nav/DirectMessageNav";
import { useLoaderData } from "react-router-dom";
import Chat from "../../chat/Chat";
import { conversationAPI } from "../../../apis/conversation.api";

const DirectMessage = () => {
  const { id, messages, users, me } = useLoaderData()
  const { username, status } = users.length === 0 ? me : users.length === 1 ? users[0] : { username: [...users, me].map(user => user.username).join(", ") }

  return (
    <div className={style.container}>
      <DirectMessageNav status={status} username={username} />
      <div className={style.content}>
        <Chat id={id} messages={messages} users={[...users, me]} key={id} createMessage={conversationAPI.newMessage} />
      </div>
    </div>
  )
}

export default DirectMessage