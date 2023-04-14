import React from "react";
import style from "./User.module.css";
import mapStatusColor from "../../helpers/mapStatusColor";

const User = ({ username, status, customStatus, avatar, color, nav, active }) => {
  const contextMenu = !nav ? "avatar" : undefined

  return (
    <div contextMenu={contextMenu} className={style.container} id={active ? style.active : ""}>
      <div contextMenu={contextMenu} className={style.image}>
        <img contextMenu={contextMenu} className={nav ? style.avatarNav : style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
        {
          !!status && <div className={nav ? style.statusNav : style.status} style={{ backgroundColor: mapStatusColor(status) }} />
        }
      </div>
      <div contextMenu={contextMenu} className={style.text}>
        <span contextMenu={contextMenu} className={nav ? style.usernameNav : style.username} style={{ color }}> {username} </span>
        <span contextMenu={contextMenu} className={nav ? style.customStatusNav : style.customStatus}> {customStatus} </span>
      </div>
    </div>
  )
}

export default User