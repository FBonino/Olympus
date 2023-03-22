import React from "react";
import { useLoaderData } from "react-router-dom";
import style from "./Server.module.css";

const Server = () => {
  const server = useLoaderData()

  return (
    <div className={style.container}>

    </div>
  )
}

export default Server