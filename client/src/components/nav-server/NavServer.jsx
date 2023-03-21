import React from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavServer.module.css";

const NavServer = ({ id, name, avatar }) => {


  return (
    <NavButton name={name}>
      <div className={style.container}>
        <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
      </div>
    </NavButton>
  )
}

export default NavServer