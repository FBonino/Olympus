import React from "react";
import NavButton from "../nav-button/NavButton";
import style from "./NavDMs.module.css";
import { AiOutlineMessage } from "react-icons/ai";
import { Link } from "react-router-dom";
import { useSelector } from "react-redux";

const NavDMs = () => {
  const { server } = useSelector(state => state.server)

  return (
    <NavButton name="Direct Messages" selected={!server}>
      <div className={style.container}>
        <Link to="/channels/@me" className={style.link}>
          <AiOutlineMessage size={24} />
        </Link>
        {
          !server && <div className={style.selected} />
        }
      </div>
    </NavButton>
  )
}

export default NavDMs