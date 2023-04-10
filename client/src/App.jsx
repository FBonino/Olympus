import React, { useEffect, useState } from 'react';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import ContextMenu from './components/context-menu/ContextMenu';
import { authAPI } from './apis/auth.api';

const App = () => {
  const [x, setX] = useState(0)
  const [y, setY] = useState(0)
  const [showMenu, setShowMenu] = useState(false)
  const [menuType, setMenuType] = useState("avatar")
  const [isLoading, setIsLoading] = useState(true)

  const closeMenu = () => setShowMenu(false)

  const handleContextMenu = e => {
    setMenuType(e.target.getAttribute("contextMenu"))
    e.preventDefault()
    setX(e.pageX)
    setY(e.pageY)
    setShowMenu(true)
  }

  useEffect(() => {
    authAPI.autoLogin()
      .then(() => setIsLoading(false))
  }, [])

  return (
    <div className="App" onContextMenu={handleContextMenu}>
      {
        !isLoading && <RouterProvider router={router} />
      }
      {
        showMenu && <ContextMenu type={menuType} top={y} left={x} handleClose={closeMenu} />
      }
    </div>
  );
}

export default App;
