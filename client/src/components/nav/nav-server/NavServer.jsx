import React, { useState } from "react";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
import NavButton from "../nav-button/NavButton";
import style from "./NavServer.module.css";

const NavServer = ({ id, name, avatar, defaultChannel }) => {
  const { server } = useSelector(state => state.server)
  const [toChannel, setToChannel] = useState(localStorage.getItem(id) || defaultChannel)

  return (
    <NavButton name={name}>
      <div className={style.container}>
        <Link to={`/channels/${id}/${toChannel}`} className={style.link}>
          <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
        </Link>
        {
          id === server?.id && <div className={style.selected} />
        }
      </div>
    </NavButton>
  )
}

export default NavServer