import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import UserSync from "./UserSync";
import HallSync from "./HallSync";
import UserHallSync from "./UserHallSync";
import ProductSync from "./ProductSync";
import { HashRouter as Router, Route, Routes, Link } from "react-router-dom";
import { Watermark } from 'antd';

const container = document.getElementById("root");
const root = createRoot(container!);

root.render( 
  <Watermark content="BETA">
  <Router>
    <div style={{ display: "flex", height: "100vh" }}>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/userSync">Sync Users</Link>
          </li>
          <li>
          <Link to="/hallSync">Sync Halls</Link>
          </li>
          <li>
          <Link to="/userHallSync">Sync User Halls</Link>
          </li>
          <li>
          <Link to="/productSync">Sync Products</Link>
          </li>
        </ul>
      </nav>
      <div style={{ flex: 1, padding: "20px", overflow: "auto" }}>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/userSync" element={<UserSync />} />
          <Route path="/hallSync" element={<HallSync />} />
          <Route path="/userHallSync" element={<UserHallSync />} />
          <Route path="/productSync" element={<ProductSync />} />
        </Routes>
      </div>
    </div>
  </Router>
  </Watermark>
);
