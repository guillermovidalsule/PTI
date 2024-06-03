import React from 'react';
import './news.css';

const News = ({ title, author, date, imageUrl, fullArticleUrl }) => {
  return (
    <div className="talos__news-container">
      <div className='talos__news-container-image'>
        <img src={imageUrl} alt={title} />
      </div>
      <div className='talos__news-container-content'>
        <h3>{title}</h3>
        <p>Author: {author}</p>
        <p>Date: {date}</p>
        <a href={fullArticleUrl} target="_blank" rel="noopener noreferrer">Read more</a>
      </div>
    </div>
  );
};

export default News;
