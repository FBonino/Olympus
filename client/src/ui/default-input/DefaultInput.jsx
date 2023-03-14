import React from "react";
import style from "./DefaultInput.module.css"

const DefaultInput = ({ type = "text", name, label, state, setState }) => {
  const handleChange = ({ target }) => setState({ ...state, [name]: target.value })

  return (
    <div className={style.container}>
      <label className={style.label}> {label} </label>
      <input className={style.input} type={type} name={name} value={state[name]} onChange={handleChange} required />
    </div>
  )
}

export default DefaultInput