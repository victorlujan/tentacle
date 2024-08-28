import {useState, useEffect} from 'react';
import './App.css';
import {GetMachines, GetUsers} from "../wailsjs/go/backend/App";
import { Skeleton , Table, TableColumnsType} from 'antd';

function App() {
interface User {
  ID: number;
  Email: string;
  Nif: string;
  Delegation: string;
}   


    
    const [machines, setMachines] = useState([]);
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState(false);
    

    function getMachines() {
        GetMachines().then((machines: any) => {
            setLoading(true);
            setMachines(machines);
            setLoading(false);
        });
    }

    function getUsers() {
        GetUsers().then((users: any) => {
            setUsers(users);
        });
    }

    useEffect(() => {
        getUsers();
    }, []);

    const columns: TableColumnsType<User> = [
        {
            title: 'ID',
            dataIndex: 'ID',
            key: 'ID',
        },
        {
            title: 'Email',
            dataIndex: 'Email',
            key: 'Email',
        },
        {
            title: 'Nif',
            dataIndex: 'Nif',
            key: 'Nif',
        },
        {
            title: 'Delegation',
            dataIndex: 'Delegation',
            key: 'Delegation',
        }
    ];


    return (
        <div id="App">
            <Table dataSource={users} columns={columns} />
        </div>
    )
}


export default App
