import React, { useEffect, useRef, useState } from "react";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import { channelAPI } from "../../../apis/channels.api";
import style from "./ChannelChat.module.css";

const ChannelChat = ({ id, messages, channelName, users }) => {
  const [updatedMessages, setUpdatedMessages] = useState(messages ? [...messages].reverse() : [])
  const lastMessageRef = useRef()
  const [message, setMessage] = useState("")
  const { id: userID } = JSON.parse(localStorage.getItem("account"))
  const { sendMessage, lastMessage } = useWebSocket(`ws://localhost:3001/api/ws/${userID}/${id}`)

  const scrollBottom = () => lastMessageRef.current?.scrollIntoView()

  const handleSubmit = async e => {
    e.preventDefault()
    const newMessage = await channelAPI.newMessage(id, message)
    sendMessage(JSON.stringify({ ...newMessage, channel: id }))
    setMessage("")
  }

  useEffect(() => scrollBottom(), [updatedMessages])

  useEffect(() => {
    if (lastMessage) {
      const { channel, ...msg } = JSON.parse(lastMessage.data)
      setUpdatedMessages(state => [...state, msg])
    }
  }, [lastMessage])

  return (
    <div className={style.container}>
      <div className={style.messages}>
        {
          updatedMessages && updatedMessages.map(m => {
            const date = new Date(m.createdAt)
            const user = users.find(u => u.id === m.author)
            return (
              <div className={style.message} key={m.id}>
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