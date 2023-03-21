import React from "react";
import style from "./NavNameModal.module.css";

const NavNameModal = ({ isOpen, name }) => {
  return isOpen && (
    <div className={style.container}>
      {name}
    </div>
  )
}

export default NavNameModal