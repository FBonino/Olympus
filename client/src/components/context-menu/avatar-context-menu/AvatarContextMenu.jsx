import React from "react";
import style from "./AvatarContextMenu.module.css";
import { MdKeyboardArrowRight } from "react-icons/md";
import ContextMenuButton from "../context-menu-button/ContextMenuButton";

const AvatarContextMenu = () => {
  return (
    <div className={style.container}>
      <ContextMenuButton text="Profile" />
      <ContextMenuButton text="Mention" />
      <ContextMenuButton text="Message" />
      <ContextMenuButton text="Call" />
      <ContextMenuButton text="Add Note" />
      <div className={style.separator} />
      <ContextMenuButton text="Invite to Server">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Add Friend" />
      <ContextMenuButton text=" Block" />
      <div className={style.separator} />
      <ContextMenuButton text="Roles">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <div className={style.separator} />
      <ContextMenuButton text="Copy ID" />
    </div>
  )
}

export default AvatarContextMenu