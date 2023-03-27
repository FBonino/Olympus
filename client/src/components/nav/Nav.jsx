import React from "react";
import { useSelector } from "react-redux";
import NavCreateServer from "./nav-create-server/NavCreateServer";
import NavDMs from "./nav-dms/NavDMs";
import NavProfile from "./nav-profile/NavProfile";
import NavServer from "./nav-server/NavServer";
import style from "./Nav.module.css";

const Nav = () => {
  const user = useSelector(state => state.user)
  const { servers } = useSelector(state => state.server)

  return (
    <div className={style.container}>
      <div className={style.subcontainer}>
        <NavDMs />
        <div className={style.separator} />
        {
          servers.map(({ id, name, avatar, defaultChannel }) =>
            <NavServer key={id} id={id} name={name} avatar={avatar} defaultChannel={defaultChannel} />
          )
        }
        <NavCreateServer />
      </div>
      <NavProfile user={user} />
    </div>
  )
}

export default Nav