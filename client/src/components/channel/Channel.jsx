import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { useLoaderData, useParams } from "react-router-dom";
import { setChannel } from "../../store/slices/server.slice";
import style from "./Channel.module.css";

const Channel = () => {
  const { id } = useParams()
  const dispatch = useDispatch()
  const channel = useLoaderData()

  useEffect(() => {
    localStorage.setItem(id, channel.id)
    dispatch(setChannel(channel))
  }, [channel, dispatch, id])

  return (
    <div className={style.container}>
      <h2> {channel.name} </h2>
    </div>
  )
}

export default Channel