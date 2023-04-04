import React from "react";
import style from "./MessageContextMenu.module.css";
import { FiLink } from "react-icons/fi";
import { HiSpeakerphone } from "react-icons/hi";
import ContextMenuButton from "../context-menu-button/ContextMenuButton";
import { MdKeyboardArrowRight, MdReply, MdOutlineNumbers, MdFlag, MdOutlineMarkChatUnread } from "react-icons/md";

const MessageContextMenu = () => {
  return (
    <div className={style.container}>
      <ContextMenuButton text="Add Reaction">
        <MdKeyboardArrowRight size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Reply">
        <MdReply size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Create Thread">
        <MdOutlineNumbers size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Mark Unread">
        <MdOutlineMarkChatUnread size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Copy Message Link">
        <FiLink size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Speak Message">
        <HiSpeakerphone size={18} />
      </ContextMenuButton>
      <ContextMenuButton text="Report Message" warning={true}>
        <MdFlag size={18} />
      </ContextMenuButton>
      <div className={style.separator} />
      <ContextMenuButton text="Copy ID" />
    </div>
  )
}

export default MessageContextMenu