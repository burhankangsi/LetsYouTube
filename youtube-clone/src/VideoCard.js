import React from 'react'
import './VideoCard.css'
import axios from 'axios';
import Avatar from "@material-ui/core/Avatar";

function VideoCard({image, title, channel, views, timestamp, channelImage}) {
    const getVideo = () => {
        axios.get('localhost:8000/{channelId}/{videoId}/video.ts')
        .then(res => {
          console.log(res.data)
        }).catch(err => {
          console.log(err)
        })
      }
  return (
    <div className="videoCard" onClick={getVideo}>
        <img className="videoCard_thumbnail" src={image} alt = ""/>
        <div className="videoCard_info">
            <Avatar 
                className="videoCard_avatar"
                alt={channel}
                src={channelImage}
            />
            <div className="videoCard_text">
                <h4>{title}</h4>
                <p>{channel}</p>
                <p>
                    {views} · {timestamp}
                </p>
            </div>
        </div>
    </div>
  );
}

export default VideoCard