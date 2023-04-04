import React from "react";
import style from "./ContextMenu.module.css";
import AvatarContextMenu from "./avatar-context-menu/AvatarContextMenu";
import ServerContextMenu from "./server-context-menu/ServerContextMenu";
import MessageContextMenu from "./message-context-menu/MessageContextMenu";
import ChannelContextMenu from "./channel-context-menu/ChannelContextMenu";

const contextMenus = {
  "message": <MessageContextMenu />,
  "avatar": <AvatarContextMenu />,
  "server": <ServerContextMenu />,
  "channel": <ChannelContextMenu />
}

const ContextMenu = ({ type, top, left, handleClose }) => {
  return type && (
    <>
      <div className={style.background} onClick={handleClose} onContextMenu={handleClose} />
      <div className={style.container} style={{ top, left }}>
        {
          contextMenus[type]
        }
      </div>
    </>
  )
}

export default ContextMenu