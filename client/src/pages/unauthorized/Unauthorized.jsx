import React from "react";
import style from "./Unauthorized.module.css";
import { useNavigate } from "react-router-dom";

const Unauthorized = () => {
  const navigate = useNavigate()

  return (
    <div className={style.container}>
      <h1> Unauthorized Access! </h1>
      <span onClick={() => navigate("/auth")}> Sign in or create an account </span>
    </div>
  )
}

export default Unauthorized