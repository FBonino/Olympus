import React, { useEffect, useState } from 'react';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { autoLogin } from './store/slices/user.slice';

const App = () => {
  const [isloading, setIsLoading] = useState(true)
  const dispatch = useDispatch()

  const handleSession = async () => {
    await dispatch(autoLogin())
    setIsLoading(false)

  }

  useEffect(() => {
    handleSession()
  })

  return (
    <div className="App">
      {
        !isloading && <RouterProvider router={router} />
      }
    </div>
  );
}

export default App;
