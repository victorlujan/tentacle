import { useState, useEffect, useRef } from "react";
import { SyncHalls } from "../wailsjs/go/backend/App";
import {
  SmileOutlined,
  CloseCircleOutlined,
} from "@ant-design/icons";
import {
  Result,
  Button,
  Progress,
  ProgressProps,
  Card,
} from "antd";

import { EventsOnMultiple } from "../wailsjs/runtime/runtime";

function HallSync() {
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<boolean>();

  const [logs, setLogs] = useState<string[]>([]);
  const [progress, setProgress] = useState<number>(0);
  const cardRef = useRef<HTMLDivElement>(null);
  const [syncButton, setSyncButton] = useState(false);

  const onImportEvent = (message: string) => {
    setLogs((log) => [...log, message]);
  };

  const onProgressEvent = (prog: number) => {
    console.log(prog);
    setProgress(prog);
  };

  EventsOnMultiple("progress", onProgressEvent, 1);
  EventsOnMultiple("hallUpdated", onImportEvent, 1);

  useEffect(() => {
    if (cardRef.current) {
      cardRef.current.scrollTop = cardRef.current.scrollHeight;
    }
  }, [logs]);

  function syncHalls() {
    setLoading(true);
    setSyncButton(true);
    SyncHalls()
      .then((status: boolean) => {
        setStatus(status);
        setLoading(false);
      })
      .finally(() => {
        setProgress(0);
        setSyncButton(false);
      });
  }

  const conicColors: ProgressProps["strokeColor"] = {
    "0%": "#87d068",
    "50%": "#ffe58f",
    "100%": "#ffccc7",
  };

  return (
    <div id="App">
      <Button onClick={syncHalls} disabled={syncButton}>
        SyncHalls
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
          title={status ? "Halls Synced" : "Error Syncing Halls"}
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
export default HallSync;
