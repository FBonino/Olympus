import React from "react";
import { Outlet } from "react-router-dom";
import Nav from "../../components/nav/Nav";
import style from "./Servers.module.css";

const Servers = () => {
  return (
    <div className={style.container}>
      <Nav />
      <Outlet />
    </div>
  )
}

export default Servers