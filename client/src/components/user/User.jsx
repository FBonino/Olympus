import React from "react";
import style from "./User.module.css";

const User = ({ username, status, customStatus, avatar, color }) => {
  const statusColor = {
    "Online": "#007000",
    "Idle": "#ceb900",
    "Do Not Disturb": "#8b0f0f",
    "Offline": "#777777",
  }

  return (
    <div className={style.container}>
      <div className={style.image}>
        <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
        <div className={style.status} style={{ backgroundColor: statusColor[status] }} />
      </div>
      <div className={style.text}>
        <span className={style.username} style={{ color }}> {username} </span>
        <span className={style.customStatus}> {customStatus} </span>
      </div>
    </div>
  )
}

export default User