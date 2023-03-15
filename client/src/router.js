import { createBrowserRouter } from "react-router-dom";
import Protected from "./components/protected/Protected";
import Auth from "./pages/auth/Auth";
import Home from "./pages/home/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Protected />,
    children: [
      {
        path: "",
        element: <Home />
      }
    ]
  },
  {
    path: "/auth",
    element: <Auth />
  }
])

export default router