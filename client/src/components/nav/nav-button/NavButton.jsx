import React, { useState } from "react";
import NavNameModal from "../nav-name-modal/NavNameModal";
import style from "./NavButton.module.css";

const NavButton = ({ children, name, selected }) => {
  const [isOpen, setIsOpen] = useState(false)

  const toggleName = () => setIsOpen(!isOpen)

  return (
    <div id={selected ? style.selected : ""} className={style.container} onMouseEnter={toggleName} onMouseLeave={toggleName}>
      {children}
      {
        isOpen && <NavNameModal name={name} />
      }
    </div>
  )
}

export default NavButton