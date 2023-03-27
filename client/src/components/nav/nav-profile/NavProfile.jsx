import React, { useState } from "react";
import User from "../../user/User";
import style from "./NavProfile.module.css";
import { AiTwotoneSetting } from "react-icons/ai";
import { FaMicrophone, FaHeadphonesAlt, FaSlash } from "react-icons/fa";

const NavProfile = ({ user }) => {
  const [state, setState] = useState({ mute: localStorage.getItem("mute") ?? false, deaf: localStorage.getItem("false") ?? false })

  const handleState = name => {
    const newState = !state[name]
    localStorage[name] = newState
    setState({ ...state, [name]: newState })
  }

  return (
    <div className={style.container}>
      <User username={user.username} status={user.status} customStatus={user.customStatus} avatar={user.avatar} />
      <div className={style.subcontainer}>
        <div className={style.iconContainer} onClick={() => handleState("mute")}>
          <FaMicrophone size={20} className={style.icon} />
          {
            state.mute && <FaSlash size={22} className={style.disabled} />
          }
        </div>
        <div className={style.iconContainer} onClick={() => handleState("deaf")}>
          <FaHeadphonesAlt size={20} className={style.icon} />
          {
            state.deaf && <FaSlash size={22} className={style.disabled} />
          }
        </div>
        <AiTwotoneSetting size={20} className={style.icon} />
      </div>
    </div>
  )
}

export default NavProfile