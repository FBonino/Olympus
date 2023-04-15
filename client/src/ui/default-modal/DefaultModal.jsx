import React from "react";
import style from "./DefaultModal.module.css";

const DefaultModal = ({ handleClose, x, y, children }) => {
  return (
    <>
      <div className={style.background} onClick={handleClose} />
      <div className={style.container} style={{ left: x + 10, top: y + 10 }}>
        {children}
      </div>
    </>
  )
}

export default DefaultModal