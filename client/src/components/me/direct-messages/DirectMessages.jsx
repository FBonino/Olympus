import React from "react";
import style from "./DirectMessages.module.css";

const DirectMessages = () => {
  return (
    <div className={style.container}>
      <form className={style.filter}>
        <input className={style.input} placeholder="Find or start a conversation" />
      </form>
      <div className={style.subcontainer}>
        <div className={style.header}>
          <span className={style.title}> DIRECT MESSAGES </span>
          <button className={style.newDM}> + </button>
        </div>
        <div className={style.dms}>

        </div>
      </div>
    </div>
  )
}

export default DirectMessages