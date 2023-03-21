import React from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavDMs.module.css";
import { AiOutlineMessage } from "react-icons/ai";

const NavDMs = () => {
  return (
    <NavButton name="Direct Messages">
      <div className={style.container}>
        <AiOutlineMessage size={24} />
      </div>
    </NavButton>
  )
}

export default NavDMs