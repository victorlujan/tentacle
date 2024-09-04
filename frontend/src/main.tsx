import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import UserSync from "./UserSync";
import HallSync from "./HallSync";
import UserHallSync from "./UserHallSync";
import ProductSync from "./ProductSync";
import {
  HashRouter as Router,
  Route,
  Routes,
  useNavigate,
} from "react-router-dom";
import {
  HomeOutlined,
  ProductOutlined,
  UserSwitchOutlined,
  UsergroupAddOutlined,
  RocketOutlined,
} from "@ant-design/icons";
import { Watermark, Menu, MenuProps } from "antd";

const container = document.getElementById("root");
const root = createRoot(container!);

type MenuItem = Required<MenuProps>["items"][number];

const items: MenuItem[] = [
  {
    key: "",
    label: "Home",
    icon: <HomeOutlined />,
  },
  {
    key: "users",
    label: "Users",
    icon: <UserSwitchOutlined />,
  },
  {
    key: "halls",
    label: "Halls",
    icon: <RocketOutlined />,
  },
  {
    key: "userhalls",
    label: "UserHalls",
    icon: <UsergroupAddOutlined />,
  },
  {
    key: "products",
    label: "Products",
    icon: <ProductOutlined />,
  },
];

const AppMenu = () => {
  const navigate = useNavigate();

  const onClick = (e: any) => {
    navigate(`/${e.key}`);
  };

  return (
    <div style={{ display: "flex", height: "100vh" }}>
      <Menu
        theme="light"
        mode="vertical"
        style={{ width: 256 }}
        items={items}
        onClick={onClick}
      />
      <div style={{ flex: 1, padding: "20px", overflow: "auto" }}>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/users" element={<UserSync />} />
          <Route path="/halls" element={<HallSync />} />
          <Route path="/userhalls" element={<UserHallSync />} />
          <Route path="/products" element={<ProductSync />} />
        </Routes>
      </div>
    </div>
  );
};

root.render(
  <Watermark content="BETA">
    <Router>
      <AppMenu />
    </Router>
  </Watermark>
);
