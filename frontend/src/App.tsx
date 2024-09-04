<<<<<<< Updated upstream
import { useState, useEffect } from "react";
import "./App.css";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { Button } from "antd";
=======
import {useState, useEffect} from 'react';
import './App.css';
import {GetMachines, GetUsers, } from "../wailsjs/go/backend/App";
import { Skeleton , Table, TableColumnsType} from 'antd';
>>>>>>> Stashed changes

function App() {
 
  return (
    <div id="App">
    </div>
  );
}

export default App;
