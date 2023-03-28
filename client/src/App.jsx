import React, { useEffect, useState } from 'react';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { autoLogin } from './store/slices/user.slice';

const App = () => {
  const dispatch = useDispatch()
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    dispatch(autoLogin())
      .then(() => setIsLoading(false))
  }, [dispatch])

  return (
    <div className="App" onContextMenu={e => e.preventDefault()}>
      {
        !isLoading && <RouterProvider router={router} />
      }
    </div>
  );
}

export default App;
