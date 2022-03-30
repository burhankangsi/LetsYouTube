import React from 'react'
import "./SearchPage.css"
import ChannelRow from './ChannelRow'
import TuneOutlinedIcon from "@material-ui/icons/TuneOutlined"
import VideoRow from './VideoRow'

function SearchPage() {
  return (
    <div className="searchPage">
        <div className="searchPage_filter">
            <TuneOutlinedIcon />
            <h2>FILTER</h2>
        </div>
        <hr/>
        <ChannelRow
          image="https://pbs.twimg.com/media/El27d6nVcAALxVX?format=png&name=large"
          channel="Mark Rober"
          verified
          subs="28M"
          noOfVideos={105}
          description="Mark Rober's engineering classes"
        />
        <hr/>

        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
        <VideoRow
          views="48M"
          subs="28M"
          description="Developed a tech setup for squirrels"
          timestamp="A few seconds ago"
          channel="Mark Rober"
          title="Making squirrels do parkour"
          image="https://i.ytimg.com/vi/hFZFjoX2cGg/maxresdefault.jpg"
        />
    </div>
  )
}

export default SearchPage;
