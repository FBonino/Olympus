import React from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavDMs.module.css";
import { AiOutlineMessage } from "react-icons/ai";
import { Link } from "react-router-dom";

const NavDMs = () => {
  return (
    <NavButton name="Direct Messages">
      <div className={style.container}>
        <Link to="/channels/@me" className={style.link}>
          <AiOutlineMessage size={24} />
        </Link>
      </div>
    </NavButton>
  )
}

export default NavDMs