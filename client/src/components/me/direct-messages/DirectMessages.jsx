import React from "react";
import style from "./DirectMessages.module.css";
import { Link } from "react-router-dom";
import User from "../../user/User";
import { useSelector } from "react-redux";

const DirectMessages = ({ conversations }) => {
  const { id: userID } = useSelector(state => state.user)

  return (
    <div className={style.container}>
      <form className={style.filter}>
        <input className={style.input} placeholder="Find or start a conversation" />
      </form>
      <div className={style.subcontainer}>
        <div className={style.header}>
          <span className={style.title}> DIRECT MESSAGES </span>
          <button className={style.newDM}> + </button>
        </div>
        <div className={style.dms}>
          {
            conversations.map(({ id, users }) => {
              const user = users.filter(u => u.id !== userID)[0]

              return (
                <Link to={`/channels/@me/${id}`} className={style.link}>
                  <User avatar={user.avatar} username={user.username} status={user.status} />
                </Link>
              )
            })
          }
        </div>
      </div>
    </div>
  )
}

export default DirectMessages