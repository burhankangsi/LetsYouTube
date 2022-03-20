import React from "react";
import "./Sidebar.css";
import SidebarRow from "./SidebarRow";
import WhatshotIcon from "@material-ui/icons/Whatshot";
import SubscriptionIcon from "@material-ui/icons/Subscriptions";
import HomeIcon from "@material-ui/icons/Home";

function Sidebar() {
  return (
    <div className="sidebar">
        <h2> I am sidebar </h2>
        <SidebarRow icon={HomeIcon} title="Home"/>
        <SidebarRow icon={WhatshotIcon} title="Trending"/>
        <SidebarRow icon={SubscriptionIcon} title="Subscription"/>
    </div>
  );
}

export default Sidebar