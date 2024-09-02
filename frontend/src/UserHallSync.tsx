import { SyncUserHalls } from "../wailsjs/go/backend/App";
import { useState, useEffect, useRef } from "react";
import { Button, Progress, ProgressProps, Result, Card } from "antd";
import { EventsOnMultiple } from "../wailsjs/runtime/runtime";
import { SmileOutlined, CloseCircleOutlined } from "@ant-design/icons";

function UserHallSync() {
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<boolean>();
  const [progress, setProgress] = useState<number>(0);
  const [syncButton, setSyncButton] = useState(false);
  const [logs, setLogs] = useState<string[]>([]);
  const cardRef = useRef<HTMLDivElement>(null);

  const onUpdateEvent = (message: string) => {
    setLogs((log) => [...log, message]);
    console.log(message);
  };

  const onProgressEvent = (prog: number) => {
    setProgress(prog);
    console.log(prog);
  };

  const conicColors: ProgressProps["strokeColor"] = {
    "0%": "#87d068",
    "50%": "#ffe58f",
    "100%": "#ffccc7",
  };

  EventsOnMultiple("progress", onProgressEvent, 1);
  EventsOnMultiple("userHallUpdated", onUpdateEvent, 1);

  useEffect(() => {
    if (cardRef.current) {
      cardRef.current.scrollTop = cardRef.current.scrollHeight;
    }
  }, [logs]);

  function syncUserHalls() {
    setLoading(true);
    setSyncButton(true);
    SyncUserHalls()
      .then((status: boolean) => {
        setStatus(status);
        setLoading(false);
      })
      .finally(() => {
        setProgress(0);
        setSyncButton(false);
      });
  }

  return (
    <div id="App">
      <Button onClick={syncUserHalls} disabled={syncButton}>
        Sync User Halls
      </Button>
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
                setStatus(false);
                setLogs([]);
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

export default UserHallSync;
