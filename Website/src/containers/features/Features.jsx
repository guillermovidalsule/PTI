import React from 'react'
import './features.css'
import {Feature} from '../../components';

const featuresData= [
  {
    title: 'Enhanced Security Coverage',
    text: 'An autonomous perimeter patrol robot offers consistent and thorough monitoring of your property. Unlike human guards, the robot can tirelessly cover the entire perimeter without breaks, reducing the risk of missed security breaches. Equipped with advanced sensors and cameras, it can detect and respond to potential threats in real-time, providing a robust and reliable security solution.'
  },
  {
    title: 'Cost-Effective Solution',
    text: 'Deploying an autonomous patrol robot can significantly lower your security costs over time. While the initial investment might seem substantial, it eliminates the recurring expenses associated with hiring and managing a team of security personnel. The robots operational efficiency and minimal maintenance needs translate to long-term savings, making it a smart financial choice for enhancing security.'
  },
  {
    title: 'Advanced Technology Integration',
    text: 'Our patrol robot integrates state-of-the-art technology, including AI-driven analytics, night vision, and thermal imaging. These features ensure that the robot can operate effectively under various conditions, identifying intrusions or anomalies that might be missed by traditional security measures. The data collected can also be used for further analysis and improvement of your security protocols.'
  }
]

const Features = () => {
  return (
    <div className='talos__features section__paddig' id="features">
      <div className='talos__features-heading'>
        <h1 className='gradient__text'>The future is now you just need to realize it. Step into future today and make it happen.</h1>
        <p>Rquest early access to get started</p>
      </div>
      <div className='talos__features-container'>
        {featuresData.map((item,index) => (
          <Feature title={item.title} text={item.text} key={item.title + index}/>
        ))}
      </div>
    </div>
  )
}

export default Features
