import React from 'react';
import {createRoot} from 'react-dom/client';
import './style.css';
import App from './App';
import Sync from './Sync';
import {BrowserRouter as Router, Route, Routes, Link} from 'react-router-dom';

const container = document.getElementById('root');
const root = createRoot(container!);

root.render(
        <Router>
        <div style={{ display: 'flex', height: '100vh' }}>
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Home</Link>
                        </li>
                        <li>
                            <Link to="/sync">Sync</Link>
                        </li>
                    </ul>
                </nav>
                <div style={{ flex: 1, padding: '20px', overflow: 'auto' }}>
                    <Routes>
                        <Route path="/" element={<App />} />
                        <Route path="/sync" element={<Sync />} />
                    </Routes>
                </div>
            </div>
        </Router>
);