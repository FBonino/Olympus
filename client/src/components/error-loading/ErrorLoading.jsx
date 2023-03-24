import React from "react";
import { useRouteError } from "react-router-dom";
import style from "./ErrorLoading.module.css";

const ErrorLoading = () => {
  const err = useRouteError()

  return (
    <div className={style.container}>
      <h1> Error loading data! </h1>
      <h3> {err.response?.data} </h3>
    </div>
  )
}

export default ErrorLoading