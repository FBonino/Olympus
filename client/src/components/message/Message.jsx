import React from "react";
import style from "./Message.module.css";

const Message = ({ avatar, username, date, content, color }) => {
  return (
    <div contextMenu="message" className={style.container}>
      <img contextMenu="avatar" className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
      <div contextMenu="message" className={style.text}>
        <div contextMenu="message" className={style.title}>
          <span contextMenu="avatar" className={style.username} style={{ color }}> {username} </span>
          <span contextMenu="message" className={style.date}> {date.toLocaleString()} </span>
        </div>
        <span contextMenu="message" className={style.content}> {content} </span>
      </div>
    </div>
  )
}

export default Message