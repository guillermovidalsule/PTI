import React from 'react';
import './blog.css';
import News from '../../components/News/News'; // Importa el componente News desde su archivo
import {image0, image1, image2,image3,image4 } from './imports';
const Blog = () => {
  return (
    <div className='talos__blog section__padding' id="blog">
      <div className='talos__blog-heading'>
        <h1 className='gradient__text'>Latest news about autonomous robots.</h1>
      </div>
      <div className='talos__blog-container'>
        <div className='talos__blog-container_groupA'>
        <News title={"AUTONOMOUS SECURITY ROBOTS ARE STARTING TO PATROL SCHOOLS AND CITIES, DESPITE CONTROVERSIES"} date="July 11, 2023" author={"Tim Mcmillan"}  imageUrl={image0} fullArticleUrl={'https://thedebrief.org/autonomous-security-robots-are-starting-to-patrol-schools-and-cities-despite-controversies/'} />
         
        </div>
        <div className='talos__blog-container_groupB'>
          <News title={"Real life Skynet? Controversial robot powered by OpenAI's ChatGPT can now have real-time conversations"} date="" author={"Julian Horsey"}  imageUrl={image1} fullArticleUrl={'https://www.geeky-gadgets.com/nvidia-humanoid-robots/'} />
          <News title={"Starship robot deliveries launches wireless on-site charging"} date="" author={"George Nott"}  imageUrl={image2} fullArticleUrl={'https://www.thegrocer.co.uk/technology-and-supply-chain/starship-robot-deliveries-launches-wireless-on-site-charging/689435.article'} />
          <News title={"Will autonomous robots replace us or create more job opportunities for humans at sea?"} date="" author={"Story by Euronews"}  imageUrl={image3} fullArticleUrl={'https://www.msn.com/en-xl/news/other/will-autonomous-robots-replace-us-or-create-more-job-opportunities-for-humans-at-sea/ar-BB1jW598'} />
          <News title={"Partnership Formed to Advance Autonomous Security Drones"} date="" author={" Glory Kaburu"}  imageUrl={image4} fullArticleUrl={'https://www.msn.com/en-us/money/news/partnership-formed-to-advance-autonomous-security-drones/ar-BB1k4NQa'} />
        </div>
      </div>
    </div>
  );
};

export default Blog;