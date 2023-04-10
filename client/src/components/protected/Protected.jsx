import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { Outlet, useLoaderData, useNavigate } from "react-router-dom";
import { setServers } from "../../store/slices/server.slice";
import { setUser } from "../../store/slices/user.slice";

const Protected = () => {
  const navigate = useNavigate()
  const dispatch = useDispatch()
  const { user, servers } = useLoaderData()

  useEffect(() => {
    if (!user) navigate("/auth")
    else {
      dispatch(setUser(user))
      dispatch(setServers(servers))
    }
  }, [dispatch, navigate, user, servers])

  return (
    <Outlet />
  )
}

export default Protected