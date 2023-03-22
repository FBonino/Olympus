import React from "react";
import { Link } from "react-router-dom";
import style from "./Home.module.css";

const Home = () => {

  return (
    <div className={style.container}>
      <Link to="/channels/@me"> Open App </Link>
      <Link> Register </Link>
    </div>
  )
}

export default Home