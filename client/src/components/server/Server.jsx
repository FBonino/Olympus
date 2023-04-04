import React, { useEffect } from "react";
import style from "./Server.module.css";
import { BiChevronDown, BiHash, BiVolumeFull } from "react-icons/bi";
import { Link, Outlet, useLoaderData, useNavigate, useParams } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { clearSelection, setServer } from "../../store/slices/server.slice";

const Server = () => {
  const dispatch = useDispatch()
  const server = useLoaderData()
  const navigate = useNavigate()
  const { id, channel } = useParams()
  const { channel: selectedChannel } = useSelector(state => state.server)

  useEffect(() => {
    const defaultChannel = localStorage.getItem(server.id) ?? server.defaultChannel
    if (id !== "@me" && !channel) navigate(`/channels/${id}/${defaultChannel}`)
  }, [server.id, server.defaultChannel, navigate, id, channel])

  useEffect(() => {
    dispatch(setServer(server))
  }, [dispatch, server])

  useEffect(() => () => dispatch(clearSelection()), [dispatch])

  return (
    <div className={style.container}>
      <div className={style.nav}>
        <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${server.avatar}`} alt="" />
        <button className={style.settings}>
          <span> {server.name} </span>
          <span> <BiChevronDown size={20} /> </span>
        </button>
        <div className={style.channels}>
          {
            server.channels?.map(c => (
              <Link contextMenu="channel" to={`/channels/${id}/${c.id}`} key={c.id} id={c.id === selectedChannel?.id ? style.selected : null} className={style.channel}>
                <span contextMenu="channel" className={style.icon}>
                  {
                    c.type === "text"
                      ? <BiHash size={21} />
                      : <BiVolumeFull size={21} />
                  }
                </span>
                <span contextMenu="channel"> {c.name} </span>
              </Link>
            ))
          }
        </div>
      </div>
      <Outlet context={server} />
    </div>
  )
}

export default Server