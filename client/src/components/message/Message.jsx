import React, { useState } from "react";
import style from "./Message.module.css";
import DefaultContextMenu from "../../ui/default-context-menu/DefaultContextMenu";

const Message = ({ avatar, username, date, content, color }) => {
  const [x, setX] = useState(0)
  const [y, setY] = useState(0)
  const [showMenu, setShowMenu] = useState(false)

  const closeMenu = () => setShowMenu(false)

  const handleContextMenu = e => {
    e.preventDefault()
    setX(e.pageX)
    setY(e.pageY)
    setShowMenu(true)
  }

  return (
    <div className={style.container} onContextMenu={handleContextMenu}>
      <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${avatar}`} alt="" />
      <div className={style.text}>
        <div className={style.title}>
          <span className={style.username} style={{ color }}> {username} </span>
          <span className={style.date}> {date.toLocaleString()} </span>
        </div>
        <span className={style.content}> {content} </span>
      </div>
      {
        showMenu && <DefaultContextMenu top={y} left={x} handleClose={closeMenu} />
      }
    </div>
  )
}

export default Message