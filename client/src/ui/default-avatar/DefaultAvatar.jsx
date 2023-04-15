import React from "react";
import style from "./DefaultAvatar.module.css";
import mapStatusColor from "../../helpers/mapStatusColor";

const DefaultAvatar = ({ avatar, status }) => {
  return (
    <div contextMenu={"avatar"} className={style.container}>
      <img contextMenu={"avatar"} className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
      {
        !!status && <div className={style.status} style={{ backgroundColor: mapStatusColor(status) }} />
      }
    </div>
  )
}

export default DefaultAvatar