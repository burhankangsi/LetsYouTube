import React from 'react'
import "./Sidebar.css";

function SidebarRow({Icon, title}) {
  return (
    <div className="sidebarRow">
        <Icon className="sidebarRow_icon"/>
        <h2 className="sidebarRow_title"> {title}</h2>
    </div>
  )
}

export default SidebarRow