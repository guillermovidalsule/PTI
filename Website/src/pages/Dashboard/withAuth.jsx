import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const withAuth = (WrappedComponent) => {
  const WithAuth = (props) => {
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
      // Aquí verificas si el usuario está autenticado
      const accessToken = localStorage.getItem('token');
      if (!accessToken) {
        // Si no hay token, redirige al usuario al inicio de sesión
        navigate('/login');
      } else {
        // Si hay token, marca al usuario como autenticado
        setIsLoggedIn(true);
      }
    }, [navigate]);

    return isLoggedIn ? <WrappedComponent {...props} /> : null;
  };

  return WithAuth;
};

export default withAuth;

