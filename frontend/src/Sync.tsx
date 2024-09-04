import {useState} from 'react';
import {SyncUsers, SyncHalls, SyncUserHalls} from "../wailsjs/go/backend/App";
import { LoadingOutlined } from '@ant-design/icons';
import { Spin } from 'antd';

function Sync(){

    const [loading, setLoading] = useState(false);
    

    function syncUsers() {
        setLoading(true);
        SyncUsers().then((status: boolean) => {
            if (status) {
                alert("Users Synced");
            } else {
                alert("Error Syncing Users");
            }
            setLoading(false);
            
        });
    }

    function syncHalls() {
        setLoading(true);
        SyncHalls().then((status: boolean) => {
            if (status) {
                alert("Halls Synced");
            } else {
                alert("Error Syncing Halls");
            }
            setLoading(false);
        });
    }

    function syncUserHalls() {
        setLoading(true);
        SyncUserHalls().then((status: boolean) => {
            if (status) {
                alert("User Halls Synced");
            } else {
                alert("Error Syncing User Halls");
            }
            setLoading(false);
        });
    }




    return (
        <div id="App">
            <button onClick={syncUsers}>Sync Users</button>
            <button onClick={syncHalls}>Sync Halls</button>
            <button onClick={syncUserHalls}>Sync User Halls</button>
            {loading ? (
                <Spin fullscreen={true} tip={"Sincronizando"} indicator={<LoadingOutlined style={{ fontSize: 100 }}  spin />} />
            ) : null}
        </div>
    )

}
export default Sync