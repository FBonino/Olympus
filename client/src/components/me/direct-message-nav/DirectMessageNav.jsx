import React from "react";
import style from "./DirectMessageNav.module.css";
import { FiAtSign } from "react-icons/fi";
import mapStatusColor from "../../../helpers/mapStatusColor";

const DirectMessageNav = ({ username, status }) => {
  return (
    <div className={style.container}>
      <FiAtSign size={22} />
      <span className={style.username}> {username} </span>
      <div className={style.status} style={{ backgroundColor: mapStatusColor(status) }} />
    </div>
  )
}

export default DirectMessageNav