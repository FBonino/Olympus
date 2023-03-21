import React from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavProfile.module.css";

const NavProfile = ({ img, status }) => {
  const statusColor = {
    "Online": "#007000",
    "Idle": "#ceb900",
    "Do Not Disturb": "#8b0f0f",
    "Offline": "#303030",
  }

  return (
    <NavButton name="Profile">
      <div className={style.container}>
        <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${img}`} alt="" />
        <div className={style.status} style={{ backgroundColor: statusColor[status] }} />
      </div>
    </NavButton>
  )
}

export default NavProfile