import React from "react";
import style from "./Message.module.css";

const Message = ({ avatar, username, date, content, color }) => {
  return (
    <div className={style.container}>
      <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
      <div className={style.text}>
        <div className={style.title}>
          <span className={style.username} style={{ color }}> {username} </span>
          <span className={style.date}> {date.toLocaleString()} </span>
        </div>
        <span className={style.content}> {content} </span>
      </div>
    </div>
  )
}

export default Message