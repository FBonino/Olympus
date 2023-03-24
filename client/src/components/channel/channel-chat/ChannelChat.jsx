import React, { useState } from "react";
import style from "./ChannelChat.module.css";

const ChannelChat = ({ messages, channelName }) => {
  const [message, setMessage] = useState()

  const handleSubmit = e => e.preventDefault()

  return (
    <div className={style.container}>
      <div className={style.messages}>

      </div>
      <form onSubmit={handleSubmit} className={style.form}>
        <input className={style.input} value={message} onChange={e => setMessage(e.target.value)} placeholder={`Message ${channelName}`} />
      </form>
    </div>
  )
}

export default ChannelChat