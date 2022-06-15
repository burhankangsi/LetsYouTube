import React from 'react'
import "./SearchPage.css"
import ChannelRow from './ChannelRow'
import TuneOutlinedIcon from "@material-ui/icons/TuneOutlined"
import VideoRow from './VideoRow'
import axios from 'axios';

const baseURL1 = "https://jsonplaceholder.typicode.com/posts/1";
const baseURL2 = "https://jsonplaceholder.typicode.com/posts/1";

function SearchPage() {
  const [post, setPost] = React.useState([]);
  const [post1, setPost1] = React.useState([]);
  let cards = []
  let channelrow = [
    {
      chanImage: '',
      channelName: '',
      subs: '',
      noofvideos: '',
      description: '',
  },
  ];
  const requestOne = axios.get(baseURL1);
  const requestTwo = axios.get(baseURL2);

  axios.all([requestOne, requestTwo]).then(axios.spread((...responses) => {
    const responseOne = responses[0]
    const responseTwo = responses[1]
    // use/access the results
    setPost(responseOne.data);
    setPost1(responseTwo.data);

    for (var i = 0; i < responseOne.data.length; i++) {
      const item = {
        id: responseOne.data[i].videoID,
        name: responseOne.data[i].videoName,
        duration: responseOne.data[i].duration,
        channelid:  responseOne.data[i].channelID,
        title:    responseOne.data[i].title,
        chanImage:  responseOne.data[i].channelImage,
        views:  responseOne.data[i].views,
        timestamp: responseOne.data[i].timestamp,
        chanName: responseOne.data[i].channelName,
        date:   responseOne.data[i].uploadDate,
        time:   responseOne.data[i].uploadTime,
        thumbnail:  responseOne.data[i].thumbnail,
      };
      cards.push(item);
    }

    channelrow.push({
      chanImage: responseTwo.data.channelImage,
      channelName: responseTwo.data.channelName,
      subs: responseTwo.data.subscribers,
      noofvideos: responseTwo.data.noOfVideos,
      description: responseTwo.data.description,
    });
     
  })).catch(errors => {
    // react on errors.
  })
  
  // React.useEffect(() => {
  //   axios.get(baseURL).then((response) => {
  //     setPost(response.data);
  //     for (var i = 0; i < response.data.length; i++) {
  //       const item = {
  //         id: response.data[i].videoID,
  //         name: response.data[i].videoName,
  //         duration: response.data[i].duration,
  //         channelid:  response.data[i].channelID,
  //         title:    response.data[i].title,
  //         chanImage:  response.data[i].channelImage,
  //         views:  response.data[i].views,
  //         timestamp: response.data[i].timestamp,
  //         chanName: response.data[i].channelName,
  //         date:   response.data[i].uploadDate,
  //         time:   response.data[i].uploadTime,
  //         thumbnail:  response.data[i].thumbnail,
  //       };
  //       cards.push(item);
  //     }
      
  //   });
  // }, 
  // []);

  if (!post) return null;
  if (!post1) return null;
  return (
    <div className="searchPage">
        <div className="searchPage_filter">
            <TuneOutlinedIcon />
            <h2>FILTER</h2>
        </div>
        <hr/>
        <ChannelRow
          image={post1.channelImage}
          channel={post1.channelName}
          verified
          subs={post1.subscribers}
          noOfVideos={post1.noOfVideos}
          description={post1.description}
        />
        {/* <ChannelRow
          image="https://pbs.twimg.com/media/El27d6nVcAALxVX?format=png&name=large"
          channel="Mark Rober"
          verified
          subs="28M"
          noOfVideos={105}
          description="Mark Rober's engineering classes"
        /> */}
        <hr/>

        {post.map((card) => {
        return (
            // <div className="post-card" key={card.id}>
            //    <h2 className="post-title">{card.title}</h2>
            //    <p className="post-body">{card.body}</p>
            //    <div className="button">
            //       <div className="delete-btn">Delete</div>
            //    </div>
            // </div>

          <VideoRow
          title={card.title}
          views={card.views}
          timestamp={card.timestamp}
          channelImage={card.channelImage}
          channel={card.channelName}
          image={card.thumbnail}
          />
        );
        })}

        {/* <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
          channelImage="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
          channelImage="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
          channelImage="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
          channelImage="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        /> */}
    </div>
  )
}

export default SearchPage;
