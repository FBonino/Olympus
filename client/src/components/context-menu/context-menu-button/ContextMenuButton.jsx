import React from "react";
import style from "./ContextMenuButton.module.css";

const ContextMenuButton = ({ text, warning, children }) => {
  return (
    <div className={`${style.container} ${warning ? style.warning : ""}`}>
      <span className={style.text}> {text} </span>
      {children}
    </div>
  )
}

export default ContextMenuButton