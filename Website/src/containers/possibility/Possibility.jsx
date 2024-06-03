import React from 'react'
import './possibility.css'
import robotimage from '../../assets/OIP.png';

const Possibility = () => {
  return (
    <div className='talos__possibility section__padding'>
      <div className='talos__possibility-image'>
        <img src= {robotimage} alt="possibility" />
      </div>
      <div className='talos__possibility-content' >
        <h4>Request early access</h4>
        <h1 className='gradient__text'>Possibilities are byond your imagination</h1>
        <p>In the not-so-distant future, surveillance robots will revolutionize security like never before. Picture a world where these sleek, autonomous guardians roam our neighborhoods, tirelessly patrolling and protecting our homes and businesses. With cutting-edge sensors and AI technology, they'll detect threats with lightning speed, responding swiftly to ensure our safety. No more worrying about false alarms or blind spots—these robots will provide comprehensive, round-the-clock surveillance, offering unparalleled peace of mind. Welcome to the future of security, where safety meets innovation!</p>
        <h4>Request early access</h4>
      </div>
    </div>
  )
}

export default Possibility
