import React, { useState } from "react";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import VideoCallIcon from "@material-ui/icons/VideoCall";
import AppsIcon from "@material-ui/icons/Apps";
import NotificationIcon from "@material-ui/icons/Notifications";
import Avatar from "@material-ui/core/Avatar";
import {Link} from "react-router-dom";
import axios from 'axios';
import './Header.css';
import UploadFile from './UploadFile';

function Header() {
  const [inputSearch, setInputSearch] = useState("");
  
  return (
    <div className="header">
      <div className="header_left">
        <MenuIcon/>
        <Link to="/">
          <img
            className="header_logo"
            src = "https://upload.wikimedia.org/wikipedia/commons/b/b8/YouTube_Logo_2017.svg"
            alt=""
            />
        </Link>
        </div>

        <div className="header_input">
          <input onChange={event => setInputSearch(event.target.value)} value={inputSearch} placeholder="Search" type="text"/>

        <Link to={`/search/${inputSearch}`}>
          <SearchIcon className="header_inputButton"/>
        </Link>
        </div>

        <div className="header_icons">
          <VideoCallIcon className="header_icon" onClick={
            <UploadFile/>
          } />
          <AppsIcon className="header_icon"/>
          <NotificationIcon className="header_icon"/>
          <Avatar
              alt="Remy Sharp"
              src="https://cdn.dribbble.com/users/1044993/screenshots/14392603/media/3af3a23806d49fb4d6585a4eded5ebc6.png?compress=1&resize=1200x900&vertical=top"
        />
        </div>
    </div>
  )
}

export default Header