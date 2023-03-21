import React, { useState } from "react";
import { useSelector } from "react-redux";
import CreateServer from "../create-server/CreateServer";
import NavDMs from "../nav-dms/NavDMs";
import NavProfile from "../nav-profile/NavProfile";
import NavServer from "../nav-server/NavServer";
import style from "./Nav.module.css";

const Nav = () => {
  const user = useSelector(state => state.user)
  // const [selectedServer, setSelectedServer] = useState()

  return (
    <div className={style.container}>
      <div className={style.subcontainer}>
        <NavDMs />
        {
          user.servers?.map(server =>
            <NavServer key={server.id} id={server.id} name={server.name} avatar={server.avatar} />
          )
        }
        <CreateServer />
      </div>
      <NavProfile img={user.avatar} status={user.status} />
    </div>
  )
}

export default Nav