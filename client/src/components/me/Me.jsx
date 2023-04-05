import React from "react";
import { Outlet, useLoaderData } from "react-router-dom";
import DirectMessages from "./direct-messages/DirectMessages";
import style from "./Me.module.css";

const Me = () => {
  const conversations = useLoaderData()

  console.log(conversations)

  return (
    <div className={style.container}>
      <DirectMessages conversations={conversations} />
      <Outlet />
    </div>
  )
}

export default Me