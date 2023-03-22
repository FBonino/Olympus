import React, { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import Nav from "../../components/nav/Nav";
import style from "./Servers.module.css";

const Servers = () => {
  // const navigate = useNavigate()

  // useEffect(() => {

  // })

  return (
    <div className={style.container}>
      <Nav />
      <Outlet />
    </div>
  )
}

export default Servers