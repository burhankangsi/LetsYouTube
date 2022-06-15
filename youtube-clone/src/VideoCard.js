import React from 'react'
import './VideoCard.css'
import axios from 'axios';
import Avatar from "@material-ui/core/Avatar";
import { useState, useCallback } from 'react'

function VideoCard({image, title, channel, views, timestamp, channelImage}) {
    // const getVideo = () => {
    //     axios.get('localhost:8000/{channelId}/{videoId}/video.ts')
    //     .then(res => {
    //       console.log(res.data)
    //     }).catch(err => {
    //       console.log(err)
    //     })
    //   }

    function play() {
      this.video.play()
    }

    function playVideo(e) {
      e.preventDefault();
      console.log('You clicked to play the video');
      getVideo(12457, 24714);
      <div>
        <video
          ref={this.getVideo}
          //src="http://techslides.com/demos/sample-videos/small.mp4"
          src={item}
          type="video/mp4"
        />
      </div>
      play();      
    }
    
    const [item, setItem] = React.useState([]);
    const getVideo = useCallback( (videoId, channelId) => {
      return async (e) => {
        e.preventDefault() //we can all this directly here now!
        axios.get('localhost:8000/{channelId}/{videoId}/video.ts')
        .then(res => {
          console.log(res.data)
          setItem(res.data)
        }).catch(err => {
          console.log(err)
        })
      }
    }, [item]);

  return (
  //  <div className="videoCard" onClick={getVideo}>
     <div className="videoCard" onClick={ {playVideo}
     }>
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