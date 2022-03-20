import React from "react";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import VideoCallIcon from "@material-ui/icons/VideoCall";
import AppsIcon from "@material-ui/icons/Apps";
import NotificationIcon from "@material-ui/icons/Notifications";
import Avatar from "@material-ui/core/Avatar";
import './Header.css';

function Header() {
  return (
    <div className="header">
      <div className="header_left">
        <MenuIcon/>
        <img
        className="header_logo" 
        src = "https://upload.wikimedia.org/wikipedia/commons/b/b8/YouTube_Logo_2017.svg"
        alt=""
        />
        </div>

        <div className="header_input">
          <input placeholder="Search" type="text"/>
          <SearchIcon className="header_inputButton"/>
        </div>

        <div className="header_icons">
          <VideoCallIcon className="header_icon"/>
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