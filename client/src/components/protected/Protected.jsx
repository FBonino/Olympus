import React, { useEffect } from "react";
import { useSelector } from "react-redux";
import { Outlet, useNavigate } from "react-router-dom";

const Protected = () => {
  const navigate = useNavigate()
  const { signedin } = useSelector(state => state.user)

  useEffect(() => {
    if (!signedin) navigate("/auth")
  })

  return (
    <Outlet />
  )
}

export default Protected