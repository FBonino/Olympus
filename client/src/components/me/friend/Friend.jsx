import React from "react";
import style from "./Friend.module.css";
import { TbMessageCircle2Filled, TbDotsVertical } from "react-icons/tb";
import mapStatusColor from "../../../helpers/mapStatusColor";
import { useNavigate } from "react-router-dom";
import { conversationAPI } from "../../../apis/conversation.api";

const Friend = ({ friend }) => {
  const navigate = useNavigate()

  const onCreateConversation = async () => {
    const conversation = await conversationAPI.create([friend.id])
    navigate(`/channels/@me/${conversation.id}`)
  }

  return (
    <div className={style.container}>
      <div className={style.user}>
        <div className={style.avatar}>
          <img className={style.image} src={`${process.env.REACT_APP_API}/uploads/${friend.avatar}`} />
          <div className={style.status} style={{ backgroundColor: mapStatusColor(friend.status) }} />
        </div>
        <div className={style.userText}>
          <span className={style.username}> {friend.username} </span>
          <span className={style.statusText}> {friend.status} </span>
        </div>
      </div>
      <div className={style.options}>
        <span className={style.option} onClick={onCreateConversation}> <TbMessageCircle2Filled className={style.icon} size={20} /> </span>
        <span className={style.option}> <TbDotsVertical className={style.icon} size={20} /> </span>
      </div>
    </div>
  )
}

export default Friend