import React from "react";
import style from "./ServerContextMenu.module.css";
import { MdKeyboardArrowRight } from "react-icons/md";
import ContextMenuButton from "../context-menu-button/ContextMenuButton";

const ServerContextMenu = () => {
  return (
    <div className={style.container}>
      <ContextMenuButton text="Mark As Read" />
      <div className={style.separator} />
      <ContextMenuButton text="Invite People" />
      <div className={style.separator} />
      <ContextMenuButton text="Mute Server" />
      <ContextMenuButton text="Notification Settings">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Hide Muted Channels" />
      <div className={style.separator} />
      <ContextMenuButton text="Privacy Settings" />
      <ContextMenuButton text="Edit Server Profile" />
      <div className={style.separator} />
      <ContextMenuButton text="Leave Server" warning={true} />
      <div className={style.separator} />
      <ContextMenuButton text="Copy ID" />
    </div>
  )
}

export default ServerContextMenu