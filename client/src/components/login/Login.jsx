import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { login } from "../../store/slices/user.slice";
import DefaultAlert from "../../ui/default-alert/DefaultAlert";
import DefaultInput from "../../ui/default-input/DefaultInput";
import style from "./Login.module.css";

const Login = ({ onChangeHasAccount }) => {
  const navigate = useNavigate()
  const dispatch = useDispatch()
  const [credentials, setCredentials] = useState({ identifier: "", password: "" })

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      await dispatch(login(credentials))
      navigate("/")
    } catch (err) {
      DefaultAlert({
        icon: "error",
        title: "Log in failed",
        text: err,
        confirmText: "Retry",
        timer: 3000
      }).then(() => setCredentials({ identifier: "", password: "" }))
    }
  }

  return (
    <div className={style.container}>
      <p className={style.title}> Welcome back! </p>
      <p className={style.subtitle}> We're so excited to see you again! </p>
      <form className={style.form} onSubmit={handleSubmit} autoComplete="off">
        <DefaultInput label="Username or Email" name="identifier" state={credentials} setState={setCredentials} />
        <DefaultInput type="password" label="Password" name="password" state={credentials} setState={setCredentials} />
        <input type="submit" className={style.submit} value="Log In" />
      </form>
      <p className={style.hasAccount} onClick={onChangeHasAccount}> Need an account? Sign up! </p>
    </div>
  )
}

export default Login