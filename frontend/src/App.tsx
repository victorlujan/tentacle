import {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {GetMachines} from "../wailsjs/go/backend/App";
import useSWR from 'swr';
import { Skeleton , Table} from 'antd';

function App() {
    const [machines, setMachines] = useState([]);
    

    function getMachines() {
        GetMachines().then((machines: any) => {
            setMachines(machines);
            console.log(machines);
        });
    }

    const columns = [
        {
            title: 'ID',
            dataIndex: 'ID',
            key: 'ID',
        },
        {
            title: 'Descripción',
            dataIndex: 'Description',
            key: 'Description',
        },
    ];



    return (
        <div id="App">
            <button onClick={getMachines}>Get Machines</button>
            <Table dataSource={machines} columns={columns} />
        </div>
    )
}


export default App
