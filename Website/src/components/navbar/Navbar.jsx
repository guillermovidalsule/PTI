import React from 'react'
import './navbar.css'
import {Link} from 'react-router-dom';
import { RiMenu3Line, RiCloseLine} from 'react-icons/ri';
import { useState } from 'react';
import logo from '../../assets/prova2.svg';

const Menu = () => (
  <>
  <p><a href='#home'> Home</a></p>
  <p><a href='#talos'> What is Talos?</a></p>
  <p><a href='#features'> Why Talos?</a></p>
  <p><a href='#about'> About Talos</a></p>
  <p><a href='#blog'> News</a></p>
  </>
)

const Navbar = () => {
  const [toggleMenu, setToggleMenu] = useState(false);
  return (
    <div className='talos_navbar'>
      <div className='talos__navbar-links'>
        <div className='talos__navbar-links_logo'>
          <img src={logo} alt="logo" />
        </div>
        <div className= 'talos__navbar-links_container'>
          <Menu />
        </div>
      </div>
      <div className='talos__navbar-sign'>
        <p><Link to='/signin'>Sign in</Link></p>
        <button type='button'><Link to='/signup'>Sign up</Link></button>
      </div>
      <div className='talos__navbar-menu'>
        {toggleMenu
          ? <RiCloseLine color='#fff' size={27} onClick={() => setToggleMenu(false)}/>
          : <RiMenu3Line color='#fff' size={27} onClick={() => setToggleMenu(true)}/>
        }
        {toggleMenu && (
          <div className='talos__navbar-menu_container scale-up-center'>
            <div className='talos__navbar-menu_container-links'> 
              <Menu />
              <div className='talos__navbar-menu_container-links-sign'>
                <p><Link to='/signin'>Sign in</Link></p>
                <button type='button'><Link to='/signup'>Sign up</Link></button>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

export default Navbar