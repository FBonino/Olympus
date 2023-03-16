import React, { useEffect } from "react";
import { useSelector } from "react-redux";
import { Outlet, useNavigate } from "react-router-dom";

const Protected = () => {
  const navigate = useNavigate()
  const user = useSelector(state => state.user)


  useEffect(() => {
    if (!user.signedin) navigate("/auth")
  })

  return (
    <Outlet />
  )
}

export default Protected