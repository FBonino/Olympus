import React, { useState } from "react";
import Login from "../../components/login/Login";
import Signup from "../../components/signup/Signup";
import style from "./Auth.module.css";

const Auth = () => {
  const [hasAccount, setHasAccount] = useState(true)

  const changeHasAccount = () => setHasAccount(!hasAccount)

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