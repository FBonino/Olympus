import React, { useState } from "react";
import { authAPI } from "../../apis/auth.api";
import DefaultAlert from "../../ui/default-alert/DefaultAlert";
import DefaultInput from "../../ui/default-input/DefaultInput";
import style from "./Signup.module.css";

const Signup = ({ onChangeHasAccount }) => {
  const [input, setInput] = useState({ username: "", email: "", password: "", passwordConfirm: "" })

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      await authAPI.signup(input)
      DefaultAlert({
        icon: "success",
        title: "Account created successfully!",
        text: "Log in to access your new account",
        confirmText: "Log in"
      }).then(() => onChangeHasAccount())
    } catch (err) {
      DefaultAlert({
        icon: "error",
        title: "Account creation failed!",
        text: err?.response?.data?.message || err?.code,
        confirmText: "Retry",
        timer: 3000
      })
    }
    setInput({ username: "", email: "", password: "", passwordConfirm: "" })
  }

  return (
    <div className={style.container}>
      <p className={style.title}> Welcome! </p>
      <p className={style.subtitle}> Create your account </p>
      <form onSubmit={handleSubmit} className={style.form} autoComplete="off">
        <DefaultInput label="Username" name="username" state={input} setState={setInput} />
        <DefaultInput label="Email" name="email" state={input} setState={setInput} />
        <DefaultInput type="password" label="Password" name="password" state={input} setState={setInput} />
        <DefaultInput type="password" label="Confirm Password" name="passwordConfirm" state={input} setState={setInput} />
        <input type="submit" className={style.submit} value="Sign up" />
      </form>
      <p className={style.hasAccount} onClick={onChangeHasAccount}> Already have an account? Log in! </p>
    </div>
  )
}

export default Signup