import React from "react";
import style from "./DefaultContextMenu.module.css";

const DefaultContextMenu = ({ top, left, handleClose, children }) => {
  return (
    <>
      <div className={style.background} onClick={handleClose} />
      <div className={style.container} style={{ top, left }}>
        {children}
      </div>
    </>
  )
}

export default DefaultContextMenu