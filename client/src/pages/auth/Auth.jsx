import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import Login from "../../components/login/Login";
import Signup from "../../components/signup/Signup";
import style from "./Auth.module.css";

const Auth = () => {
  const navigate = useNavigate()
  const user = useSelector(state => state.user)
  const [hasAccount, setHasAccount] = useState(true)

  const changeHasAccount = () => setHasAccount(!hasAccount)

  useEffect(() => {
    if (user) navigate("/channels/@me")
  }, [user, navigate])

  return (
    <div className={style.container}>
      {
        hasAccount
          ? <Login onChangeHasAccount={changeHasAccount} />
          : <Signup onChangeHasAccount={changeHasAccount} />
      }
    </div>
  )
}

export default Auth