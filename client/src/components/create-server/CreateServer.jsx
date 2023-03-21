import React from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./CreateServer.module.css";
import { BiPlus } from "react-icons/bi";

const CreateServer = () => {
  return (
    <NavButton name="Create a Server">
      <div className={style.container}>
        <BiPlus size={28} />
      </div>
    </NavButton>
  )
}

export default CreateServer