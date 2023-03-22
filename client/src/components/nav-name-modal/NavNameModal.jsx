import React from "react";
import style from "./NavNameModal.module.css";

const NavNameModal = ({ name }) => {
  return (
    <div className={style.container}>
      {name}
    </div>
  )
}

export default NavNameModal