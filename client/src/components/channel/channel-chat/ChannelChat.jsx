import React, { useEffect, useRef, useState } from "react";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import style from "./ChannelChat.module.css";
import Message from "../../message/Message";

const ChannelChat = ({ id, messages, channelName, users, roles, createMessage }) => {
  const [updatedMessages, setUpdatedMessages] = useState(messages ? [...messages].reverse() : [])
  const lastMessageRef = useRef()
  const [message, setMessage] = useState("")
  const { id: userID } = JSON.parse(localStorage.getItem("account"))
  const { sendMessage, lastMessage } = useWebSocket(`ws://localhost:3001/api/ws/${userID}/${id}`)

  const scrollBottom = () => lastMessageRef.current?.scrollIntoView()

  const handleSubmit = async e => {
    e.preventDefault()
    const newMessage = await createMessage(id, message)
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
            const role = roles?.find(r => r.id === user.roles[0])
            return (
              <Message key={m.id} avatar={user.avatar} username={user.username} date={date} content={m.content} color={role?.color} />
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