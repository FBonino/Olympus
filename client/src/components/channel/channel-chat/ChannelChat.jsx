import React, { useEffect, useRef, useState } from "react";
import { channelAPI } from "../../../apis/channels.api";
import style from "./ChannelChat.module.css";

const ChannelChat = ({ id, messages, channelName, users }) => {
  const lastMessageRef = useRef()
  const [message, setMessage] = useState("")

  const scrollBottom = () => lastMessageRef.current?.scrollIntoView()

  const handleSubmit = async e => {
    e.preventDefault()
    const data = await channelAPI.newMessage(id, message)
    setMessage("")
  }

  useEffect(() => scrollBottom(), [messages])

  return (
    <div className={style.container}>
      <div className={style.messages}>
        {
          messages && messages.map(m => {
            const date = new Date(m.createdAt)
            const user = users.find(u => u.id === m.author)
            return (
              <div className={style.message}>
                <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${user.avatar}`} alt="" />
                <div className={style.text}>
                  <div className={style.title}>
                    <span className={style.username}> {user.username} </span>
                    <span className={style.date}> {date.toLocaleString()} </span>
                  </div>
                  <span className={style.content}> {m.content} </span>
                </div>
              </div>
            )
          })
        }
        <div ref={lastMessageRef} />
      </div>
      <form onSubmit={handleSubmit} className={style.form}>
        <input className={style.input} value={message} onChange={e => setMessage(e.target.value)} placeholder={`Message ${channelName}`} />
      </form>
    </div>
  )
}

export default ChannelChat