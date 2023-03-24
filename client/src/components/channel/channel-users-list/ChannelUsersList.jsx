import React from "react";
import style from "./ChannelUsersList.module.css";

const ChannelUsersList = ({ users }) => {
  return (
    <div className={style.container}>
      {
        users && users.map(u => (
          <div className={style.user} key={u.id}>
            <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${u.avatar}`} alt="" />
            <div className={style.text}>
              <span className={style.username}> {u.username} </span>
              <span className={style.customStatus}> {u.customStatus} </span>
            </div>
          </div>
        ))
      }
    </div>
  )
}

export default ChannelUsersList