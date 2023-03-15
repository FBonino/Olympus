import React from "react";
import style from "./NavButton.module.css";

const NavButton = ({ img, type, status, name, handleClick }) => {
  const statusColor = {
    "Online": "#007000",
    "Idle": "#ceb900",
    "Do Not Disturb": "#8b0f0f",
    "Offline": "#303030",
  }

  return (
    <div className={style.container}>
      <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${img}`} alt="" />
      {
        type === "profile" && <div className={style.status} style={{ backgroundColor: statusColor[status] }} />
      }
    </div>
  )
}

export default NavButton