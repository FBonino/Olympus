import React from "react";
import { Link } from "react-router-dom";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavServer.module.css";

const NavServer = ({ id, name, avatar }) => {


  return (
    <NavButton name={name}>
      <div className={style.container}>
        <Link to={`/channels/${id}`} className={style.link}>
          <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
        </Link>
      </div>
    </NavButton>
  )
}

export default NavServer