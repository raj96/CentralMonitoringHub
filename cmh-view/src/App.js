import { Build, Home as HomeIcon, Settings, Source } from "@mui/icons-material";
import { Container } from "@mui/material";
import { useState } from "react";
import { HashRouter as Router, Route, Routes, Navigate } from "react-router-dom";
import "./App.css";

import Navbar from "./components/Navbar";
import SideMenu from "./components/SideMenu";
import Home from "./components/Home";
import SourceSettings from "./components/SourceSettings";


function App() {
  const validPages = { "Source Settings": <SourceSettings /> };
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [currentPageTag, setCurrentPageTag] = useState("Home");
  const [containerMaxWidth, setContainerMaxWidth] = useState("xl");
  const toggleDrawer = () => {
    setDrawerOpen((prevDrawerState) => !prevDrawerState);
  };
  const sideMenuContent = [
    {
      icon: <HomeIcon />,
      text: "Home",
      parentContainerMaxWidth: "xl",
      location: "/home",
    },
    {
      icon: <Build />,
      text: "Triggers",
      parentContainerMaxWidth: "xl",
      location: "/triggers",
    },
    {
      icon: <Source />,
      text: "Source Settings",
      parentContainerMaxWidth: "md",
      location: "/sourcesettings",
    },
    {
      icon: <Settings />,
      text: "Account Settings",
      parentContainerMaxWidth: "xl",
      location: "/accountsettings",
    }
  ];
  const sideMenuCallback = (clickedMenuItemName) => {
    if (validPages[clickedMenuItemName]) {
      setCurrentPageTag(clickedMenuItemName);
    }
  };

  return (
    <div className="App">
      <Router>
        <Navbar menuClicked={toggleDrawer} />
        <SideMenu
          drawerOpen={drawerOpen}
          setDrawerOpen={setDrawerOpen}
          menuContent={sideMenuContent}
          menuSelectCallback={sideMenuCallback}
          setParentContainerMaxWidth={setContainerMaxWidth}
        />
        <Container maxWidth={containerMaxWidth}>
          <Routes>
            <Route path="/sourcesettings" element={<SourceSettings />} />
            <Route path="/home" element={<Home />} />
            <Route path="/" element={<Navigate to="/home" />} />
          </Routes>
        </Container>
      </Router>
    </div>
  );
}

export default App;
