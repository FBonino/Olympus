import React from "react";
import { Outlet } from "react-router-dom";
import DirectMessages from "./direct-messages/DirectMessages";
import style from "./Me.module.css";

const Me = () => {
  return (
    <div className={style.container}>
      <DirectMessages />
      <Outlet />
    </div>
  )
}

export default Me