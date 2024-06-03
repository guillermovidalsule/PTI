import React, { useState } from 'react';
import './signup.css';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import { FaUser, FaLock, FaEnvelope } from "react-icons/fa";

const Signup = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [repeatPassword, setRepeatPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (password !== repeatPassword) {
      setError('Passwords do not match');
      return;
    }

    try {
      const xhr = new XMLHttpRequest();
      xhr.open("POST", "http://nattech.fib.upc.edu:40342/users/signup", true);
      xhr.setRequestHeader("Content-Type", "application/json");

      xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
          if (xhr.status === 200) {
            const data = JSON.parse(xhr.responseText);
            console.log(data);
            const token = data.token; // Suponiendo que el token se devuelve en la propiedad "token" del objeto de respuesta
            localStorage.setItem('token', token);
            toast.success('Registered successfully.');
            navigate('/dash');
          } else {
            setError('User already registred');
          }
        }
      };

      xhr.onerror = function () {
        toast.error('Failed to make request.');
      };

      xhr.send(JSON.stringify({ email, username, password }));
    } catch (error) {
      toast.error('Failed: ' + error.message);
    }
  };

  return (
    <div className='total' >
      <div className='signup'>
        <form action='' onSubmit={handleSubmit}>
          <h1>Sign Up</h1>

          <div className='input-box'>
            <input type='email' placeholder='Email'
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required />
            <FaEnvelope className='icon' />
          </div>

          <div className='input-box'>
            <input type='text' placeholder='Username'
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required />
            <FaUser className='icon' />
          </div>

          <div className='input-box'>
            <input type='password' placeholder='Password'
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required />
            <FaLock className='icon' />
          </div>

          <div className='input-box'>
            <input type='password' placeholder='Repeat Password'
              value={repeatPassword}
              onChange={(e) => setRepeatPassword(e.target.value)}
              required />
            <FaLock className='icon' />
          </div>

          {error && <p className='error'>{error}</p>}
          <button type='submit'>Sign Up</button>
          <div className='register-link'>
            <p>I already have an account <a href='/signin'>Sign in</a></p>
          </div>

        </form>
      </div>
    </div>
  );
}

export default Signup;
