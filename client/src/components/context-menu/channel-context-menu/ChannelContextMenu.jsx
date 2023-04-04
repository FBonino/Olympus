import React from "react";
import style from "./ChannelContextMenu.module.css";
import { MdKeyboardArrowRight } from "react-icons/md";
import ContextMenuButton from "../context-menu-button/ContextMenuButton";

const ChannelContextMenu = () => {
  return (
    <div className={style.container}>
      <ContextMenuButton text="Mark As Read" />
      <div className={style.separator} />
      <ContextMenuButton text="Invite People" />
      <ContextMenuButton text="Copy Link" />
      <div className={style.separator} />
      <ContextMenuButton text="Mute Channel">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Notification Settings">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <div className={style.separator} />
      <ContextMenuButton text="Edit Channel" />
      <ContextMenuButton text="Duplicate Channel" />
      <ContextMenuButton text="Create Text Channel" />
      <ContextMenuButton text="Delete Channel" warning={true} />
      <div className={style.separator} />
      <ContextMenuButton text="Copy ID" />
    </div>
  )
}

export default ChannelContextMenu