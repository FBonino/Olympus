import axios from "axios";

const api = axios.create({
  baseURL: `${process.env.REACT_APP_API}/api`
})

const errorHandler = async error => {
  const statusCode = error.response?.status

  if (error.code === "ERR_CANCELED") {
    return Promise.resolve()
  }

  if (statusCode && statusCode !== 401) {
    console.error(error)
  }

  return Promise.reject(error)
}

api.interceptors.response.use(undefined, async error => {
  return await errorHandler(error)
})

export default api