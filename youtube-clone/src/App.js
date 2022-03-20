// import logo from './logo.svg';
import React from "react";
import Header from "./Header";
import Sidebar from "./Sidebar";

import './App.css';
import RecommendedVideos from "./RecommendedVideos";

function App() {
  return (
    <div className="App">
      <Header/>
      <div className="app_page">
        <Sidebar/>
        <RecommendedVideos/>
      </div>
    </div>
  );
}

export default App;
