import React from "react";
import style from "./FriendsNav.module.css";

const FriendsNav = ({ selected, setSelected, resetSearch }) => {
  const onSelect = ({ target }) => {
    resetSearch()
    setSelected(target.value)
  }

  return (
    <div className={style.container}>
      <span className={style.title}> Friends </span>
      <div className={style.separator} />
      <button value="Online" className={style.button} id={selected === "Online" ? style.selected : ""} onClick={onSelect}> Online </button>
      <button value="All friends" className={style.button} id={selected === "All friends" ? style.selected : ""} onClick={onSelect}> All </button>
      <button value="Pending" className={style.button} id={selected === "Pending" ? style.selected : ""} onClick={onSelect}> Pending </button>
      <button value="Blocked" className={style.button} id={selected === "Blocked" ? style.selected : ""} onClick={onSelect}> Blocked </button>
      <button className={style.new}> Add Friend </button>
    </div>
  )
}

export default FriendsNav