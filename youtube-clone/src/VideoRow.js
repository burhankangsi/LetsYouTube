import React from 'react'
import './VideoRow.css'
import axios from 'axios';
import { useState } from 'react';

function VideoRow({views, subs, description, timestamp, channel, title, image}) {
//  const [video] = useState('')
  const getVideo = () => {
    axios.get('localhost:8000/{channelId}/{videoId}/video.ts')
    .then(res => {
      console.log(res.data)
    }).catch(err => {
      console.log(err)
    })
  }
  return (
    <div className="videoRow" onClick={getVideo}>
        <img src={image} alt=""/>
        <div className="videoRow_text">
            <h3> {title} </h3>
            <p className="videoRow_headline">
                {channel} . {" "} 
                <span className="videoRow_subs">
                    <span className="videoRow_subsNumber"> {subs} </span> Subscribers
                </span>{" "}
                {views} views . {timestamp}
            </p>
            <p className="videoRow_description">
                 {description} </p>
        </div>
    </div>
  )
}

export default VideoRow