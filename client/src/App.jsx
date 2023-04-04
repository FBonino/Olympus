import React, { useEffect, useState } from 'react';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { autoLogin } from './store/slices/user.slice';
import ContextMenu from './components/context-menu/ContextMenu';

const App = () => {
  const dispatch = useDispatch()
  const [isLoading, setIsLoading] = useState(true)
  const [x, setX] = useState(0)
  const [y, setY] = useState(0)
  const [showMenu, setShowMenu] = useState(false)
  const [menuType, setMenuType] = useState("avatar")

  const closeMenu = () => setShowMenu(false)

  const handleContextMenu = e => {
    setMenuType(e.target.getAttribute("contextMenu"))
    e.preventDefault()
    setX(e.pageX)
    setY(e.pageY)
    setShowMenu(true)
  }

  useEffect(() => {
    dispatch(autoLogin())
      .then(() => setIsLoading(false))
  }, [dispatch])

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
