import React from 'react'
import "./LikedVideos.css";
import VideoCard from './VideoCard';
import axios from 'axios';

const baseURL = "https://flash-api.herokuapp.com/feed/libray/likedvideos/{channelId}";

function LikedVideos() {
  const [post, setPost] = React.useState([]);
  let cards = []
  
  React.useEffect(() => {
    axios.get(baseURL).then((response) => {
      setPost(response.data);
      for (var i = 0; i < response.data.length; i++) {
        const item = {
          id: response.data[i].videoID,
          name: response.data[i].videoName,
          duration: response.data[i].duration,
          channelid:  response.data[i].channelID,
          title:    response.data[i].title,
          chanImage:  response.data[i].channelImage,
          views:  response.data[i].views,
          timestamp: response.data[i].timestamp,
          chanName: response.data[i].channelName,
          date:   response.data[i].uploadDate,
          time:   response.data[i].uploadTime,
          thumbnail:  response.data[i].thumbnail,
        };
        cards.push(item);
      }
      
    });
  }, 
  []);

  if (!post) return null;

  return (
    <div className="likedVideos">
      <h2>Liked Videos</h2>
      <div className="liked_videos">

      {post.map((card) => {
       return (
          // <div className="post-card" key={card.id}>
          //    <h2 className="post-title">{card.title}</h2>
          //    <p className="post-body">{card.body}</p>
          //    <div className="button">
          //       <div className="delete-btn">Delete</div>
          //    </div>
          // </div>

        <VideoCard
        title={card.title}
        views={card.views}
        timestamp={card.timestamp}
        channelImage={card.channelImage}
        channel={card.channelName}
        image={card.thumbnail}
        />
       );
    })}
                
        {/* <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        />
        <VideoCard
        title="How to learn coding from scratch"
        views="272k Views"
        timestamp="24 days ago"
        channelImage="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRz3SPbs9loLmVpdTIEYQ7Lb2XUbOR7Uz84zg&usqp=CAU"
        channel="Code camp"
        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQrWxCScou_U4K_XxQsg2b-fuSwCQy11JDdgg&usqp=CAU"
        />
        <VideoCard
        title="Last one to take hands off this lamborghini keeps it"
        views="48M Views"
        timestamp="2 days ago"
        channelImage="https://logos-world.net/wp-content/uploads/2021/09/Mr-Beast-Logo.png"
        channel="MrBeast"
        image="https://i.ytimg.com/vi/KSKJKLmAqpI/mqdefault.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://pbs.twimg.com/media/El27d6nVcAALxVX?format=png&name=large"
        channel="Mark Rober"
        image="https://i.ytimg.com/vi/KSKJKLmAqpI/mqdefault.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        />
        <VideoCard
        title="How to learn coding from scratch"
        views="272k Views"
        timestamp="24 days ago"
        channelImage="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRz3SPbs9loLmVpdTIEYQ7Lb2XUbOR7Uz84zg&usqp=CAU"
        channel="Code camp"
        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQrWxCScou_U4K_XxQsg2b-fuSwCQy11JDdgg&usqp=CAU"
        />
        <VideoCard
        title="Last one to take hands off this lamborghini keeps it"
        views="48M Views"
        timestamp="2 days ago"
        channelImage="https://logos-world.net/wp-content/uploads/2021/09/Mr-Beast-Logo.png"
        channel="MrBeast"
        image="https://i.ytimg.com/vi/KSKJKLmAqpI/mqdefault.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://pbs.twimg.com/media/El27d6nVcAALxVX?format=png&name=large"
        channel="Mark Rober"
        image="https://i.ytimg.com/vi/KSKJKLmAqpI/mqdefault.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        />
        <VideoCard
        title="Become a UI developer in 10 days"
        views="1.4M Views"
        timestamp="3 days ago"
        channelImage="https://cdn4.vectorstock.com/i/1000x1000/20/78/ninja-sport-mascot-logo-design-vector-26472078.jpg"
        channel="UI Ninja"
        image="https://www.learntek.org/wp-content/uploads/2017/09/UI-DEVeloper-1.jpg"
        /> */}
      </div>
    </div>
  )
}

export default LikedVideos