import React from "react";
import style from "./DefaultModal.module.css";
import { AiOutlineCloseCircle } from "react-icons/ai";

const DefaultModal = ({ handleClose, children }) => {
  return (
    <>
      <div className={style.background} onClick={handleClose} />
      <div className={style.container}>
        <button onClick={handleClose} className={style.close}> <AiOutlineCloseCircle size={28} /> </button>
        {children}
      </div>
    </>
  )
}

export default DefaultModal