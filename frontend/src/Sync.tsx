import { useState, useEffect, useRef } from "react";
import { SyncUsers, SyncHalls, SyncUserHalls } from "../wailsjs/go/backend/App";
import {
  LoadingOutlined,
  SmileOutlined,
  CloseCircleOutlined,
} from "@ant-design/icons";
import {
  Spin,
  Result,
  Button,
  Progress,
  ProgressProps,
  Card,
  Flex,
} from "antd";

import { EventsOnMultiple } from "../wailsjs/runtime/runtime";

function Sync() {
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<boolean>();

  const [logs, setLogs] = useState<string[]>([]);
  const [progress, setProgress] = useState<number>(0);
  const cardRef = useRef<HTMLDivElement>(null);
  const [userButton, setUserButton] = useState(false);

  const onImportEvent = (message: string) => {
    setLogs((log) => [...log, message]);
    console.log(message);
  };

  const onProgressEvent = (prog: number) => {
    setProgress(prog);
  };

  EventsOnMultiple("progress", onProgressEvent, 1);
  EventsOnMultiple("userUpdated", onImportEvent, 1);

  // useEffect(() => {
  // }, [logs]);

  useEffect(() => {
    if (cardRef.current) {
      cardRef.current.scrollTop = cardRef.current.scrollHeight;
    }
  }, [logs]);

  function syncUsers() {
    setLoading(true);
    setUserButton(true);
    SyncUsers()
      .then((status: boolean) => {
        setStatus(status);
        setLoading(false);
      })
      .finally(() => {
        setProgress(0);
        setUserButton(false);
      });
  }

  function syncHalls() {
    setLoading(true);
    SyncHalls().then((status: boolean) => {
      setLoading(false);
    });
  }

  function syncUserHalls() {
    setLoading(true);
    SyncUserHalls()
      .then((status: boolean) => {
        setStatus(status);
        setLoading(false);
      })
      .finally(() => {
        setProgress(0);
      });
  }

  const conicColors: ProgressProps["strokeColor"] = {
    "0%": "#87d068",
    "50%": "#ffe58f",
    "100%": "#ffccc7",
  };

  return (
    <div id="App">
      <Button onClick={syncUsers} disabled={userButton}>
        Sync Users
      </Button>
      <Button onClick={syncHalls}>Sync Halls</Button>
      <Button onClick={syncUserHalls}>Sync User Halls</Button>
      {loading ? (
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            flexDirection: "column",
            marginTop: 20,
          }}
        >
          <Progress
            type="dashboard"
            percent={Number(progress.toFixed(1))}
            strokeColor={conicColors}
          />
        </div>
      ) : null}

      {status ? (
        <Result
          status={status ? "success" : "error"}
          title={status ? "User Halls Synced" : "Error Syncing User Halls"}
          icon={
            status ? (
              <SmileOutlined style={{ fontSize: 100 }} />
            ) : (
              <CloseCircleOutlined style={{ fontSize: 100 }} />
            )
          }
          extra={[
            <Button
              type="primary"
              key="console"
              onClick={() => {
                setStatus(false), setLogs([]);
              }}
            >
              Go Home
            </Button>,
          ]}
        />
      ) : null}

      {logs.length > 0 && (
        <Card
          title="Logs"
          style={{ marginTop: 20, height: 400, overflow: "auto" }}
          styles={{
            header: {
              position: "sticky",
              top: 0,
              zIndex: 1,
              background: "#fff",
            },
          }}
          ref={cardRef}
        >
          {logs.map((log, index) => (
            <p key={index}>{log}</p>
          ))}
        </Card>
      )}
    </div>
  );
}
export default Sync;
