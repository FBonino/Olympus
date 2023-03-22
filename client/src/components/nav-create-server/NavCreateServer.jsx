import React, { useState } from "react";
import NavButton from "../../ui/nav-button/NavButton";
import style from "./NavCreateServer.module.css";
import { BiPlus } from "react-icons/bi";
import CreateServerForm from "../create-server-form/CreateServerForm";

const NavCreateServer = () => {
  const [isOpen, setIsOpen] = useState(false)

  const toggleModal = () => setIsOpen(!isOpen)

  return (
    <NavButton name="Create a Server">
      <div className={style.container} onClick={toggleModal}>
        <BiPlus size={28} />
      </div>
      {
        isOpen && <CreateServerForm handleClose={toggleModal} />
      }
    </NavButton>
  )
}

export default NavCreateServer