import React from 'react'
import './header.css'
import {Link} from 'react-router-dom';
import people from '../../assets/people.png';
import robot from '../../assets/protecting-robot-and-shield-security-3d-illustration-isolated-HW162N-removebg-preview.png';



const Header = () => {
  return (
    <div className='talos__header section__padding' id="home">
      <div className='talos__header-content'>
        <h1 className='gradient__text'>Let's build something amazing with Talos</h1>
        <p>Yet bed any for travelling assistance indulgence unpleasing. Not thoughts all exercise blessing. Indulgence way everything joy alteration boisterous the attachment. Party we years to order allow asked of.</p>
        <div className='talos__header-content__input'>
          <input type='email' placeholder='Your Email Address'></input>
          <Link to="/signup">
            <button type='button' >Get Started</button>
          </Link>
        </div>
        <div className='talos__header-content__people'>
          <img src={people} alt="people"/>
          <p>2.000 people requested access a visit in last 24 hours</p>
        </div>
      </div>
      <div className='talos__header-image'>
          <img src={robot} alt="robot"/>
      </div> 
    </div>
  )
}

export default Header
