import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { Outlet, useLoaderData } from "react-router-dom";
import { setServers } from "../../store/slices/server.slice";
import { setUser } from "../../store/slices/user.slice";

const Protected = () => {
  const dispatch = useDispatch()
  const { user, servers } = useLoaderData()

  useEffect(() => {
    dispatch(setUser(user))
    dispatch(setServers(servers))
  }, [dispatch, user, servers])

  return (
    <Outlet />
  )
}

export default Protected