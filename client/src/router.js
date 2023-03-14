import { createBrowserRouter } from "react-router-dom";
import Auth from "./pages/auth/Auth";
import Home from "./pages/home/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />
  },
  {
    path: "/auth",
    element: <Auth />
  }
])

export default router