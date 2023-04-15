import React, { useState } from "react";
import style from "./DirectMessages.module.css";
import { Link, useParams } from "react-router-dom";
import User from "../../user/User";
import { HiPlusSm } from "react-icons/hi";
import { useSelector } from "react-redux";
import CreateConversationForm from "../../create-conversation-form/CreateConversationForm";

const DirectMessages = ({ conversations }) => {
  const params = useParams()
  const [isOpen, setIsOpen] = useState(false)
  const { friends } = useSelector(state => state.user)
  const [coords, setCoords] = useState({ x: 0, y: 0 })

  const openModal = e => {
    setCoords({ x: e.pageX, y: e.pageY })
    setIsOpen(true)
  }

  const closeModal = e => {
    setIsOpen(false)
  }

  return (
    <div className={style.container}>
      <form className={style.filter}>
        <input className={style.input} placeholder="Find or start a conversation" />
      </form>
      <div className={style.subcontainer}>
        <div className={style.header}>
          <span className={style.title}> DIRECT MESSAGES </span>
          <button className={style.newDM} onClick={openModal}> <HiPlusSm size={20} /> </button>
        </div>
        <div className={style.dms}>
          {
            !!conversations.length && conversations.map(({ id, users, avatar, me }) => {
              const user = users.length === 0 ? me : users.length === 1 ? users[0] : { avatar, username: [...users, me].map(user => user.username).join(", ") }
              return (
                <Link key={id} to={`/channels/@me/${id}`} className={style.link}>
                  <div className={style.wrapper}>
                    <User avatar={user.avatar} username={user.username} status={user.status} active={id === params.conversation} />
                  </div>
                </Link>
              )
            })
          }
        </div>
      </div>
      {
        isOpen && <CreateConversationForm friends={friends.filter(f => f.relation === "Friend")} handleClose={closeModal} coords={coords} />
      }
    </div>
  )
}

export default DirectMessages