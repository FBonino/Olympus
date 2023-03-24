import React from "react";
import style from "./ChannelNav.module.css";
import { BiHash, BiVolumeFull } from "react-icons/bi";

const ChannelNav = ({ name, type, topic }) => {

  return (
    <div className={style.container}>
      <span className={style.icon}>
        {
          type === "text"
            ? <BiHash size={22} />
            : <BiVolumeFull size={22} />
        }
      </span>
      <span> {name} </span>
      <span> {topic} </span>
    </div>
  )
}

export default ChannelNav