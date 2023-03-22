import React, { useState } from "react";
import NavNameModal from "../../components/nav-name-modal/NavNameModal";
import style from "./NavButton.module.css";

const NavButton = ({ children, name }) => {
  const [isOpen, setIsOpen] = useState(false)

  const toggleName = () => setIsOpen(!isOpen)

  return (
    <div className={style.container} onMouseEnter={toggleName} onMouseLeave={toggleName}>
      {children}
      {
        isOpen && <NavNameModal name={name} />
      }
    </div>
  )
}

export default NavButton