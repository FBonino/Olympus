import React from "react";
import { useSelector } from "react-redux";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./Nav.module.css";

const Nav = () => {
  const user = useSelector(state => state.user)

  return (
    <div className={style.container}>
      <div className={style.subcontainer}>
        {
          user.servers?.map(server =>
            <NavButton img={server.avatar} type="server" name={server.name} key={server.id} />
          )
        }
      </div>
      <NavButton img={user.avatar} type="profile" status={user.status} name={user.username} />
    </div>
  )
}

export default Nav