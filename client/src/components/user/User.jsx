import React from "react";
import style from "./User.module.css";

const User = ({ username, status, customStatus, avatar, color, nav }) => {
  const statusColor = {
    "Online": "#007000",
    "Idle": "#ceb900",
    "Do Not Disturb": "#8b0f0f",
    "Offline": "#777777",
  }

  const contextMenu = !nav ? "avatar" : undefined

  return (
    <div contextMenu={contextMenu} className={style.container}>
      <div contextMenu={contextMenu} className={style.image}>
        <img contextMenu={contextMenu} className={nav ? style.avatarNav : style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
        <div className={nav ? style.statusNav : style.status} style={{ backgroundColor: statusColor[status] }} />
      </div>
      <div contextMenu={contextMenu} className={style.text}>
        <span contextMenu={contextMenu} className={nav ? style.usernameNav : style.username} style={{ color }}> {username} </span>
        <span contextMenu={contextMenu} className={nav ? style.customStatusNav : style.customStatus}> {customStatus} </span>
      </div>
    </div>
  )
}

export default User