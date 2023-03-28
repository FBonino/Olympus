import React from "react";
import FriendsNav from "../friends-nav/FriendsNav";
import style from "./Friends.module.css";

const Friends = () => {

  return (
    <div className={style.container}>
      <FriendsNav />
    </div>
  )
}

export default Friends