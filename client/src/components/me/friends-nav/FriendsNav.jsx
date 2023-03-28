import React, { useState } from "react";
import style from "./FriendsNav.module.css";

const FriendsNav = () => {
  const [selected, setSelected] = useState("online")

  const onSelect = ({ target }) => setSelected(target.value)

  return (
    <div className={style.container}>
      <span className={style.title}> Friends </span>
      <div className={style.separator} />
      <button value="online" className={style.button} id={selected === "online" ? style.selected : ""} onClick={onSelect}> Online </button>
      <button value="all" className={style.button} id={selected === "all" ? style.selected : ""} onClick={onSelect}> All </button>
      <button value="pending" className={style.button} id={selected === "pending" ? style.selected : ""} onClick={onSelect}> Pending </button>
      <button value="blocked" className={style.button} id={selected === "blocked" ? style.selected : ""} onClick={onSelect}> Blocked </button>
      <button className={style.new}> Add Friend </button>
    </div>
  )
}

export default FriendsNav