import React from "react";
import { Layout, Menu } from "antd";
import { Link } from "react-router-dom";
import ThemeSwitcher from "./ThemeSwitcher";

const { Header } = Layout;

const TopMenu = () => {
  return (
    <Layout>
      <Header className="flex justify-between items-center">
        <div className="flex justify-center flex-grow">
          <Menu className="flex-grow justify-center" theme="dark" mode="horizontal" defaultSelectedKeys={["1"]}>
            <Menu.Item key="1">
              <Link to="/">Home</Link>
            </Menu.Item>
            <Menu.Item key="2">
              <Link to="/life-circle">Life Circle</Link>
            </Menu.Item>
            <Menu.Item key="3">
              <Link to="/profile">Profile</Link>
            </Menu.Item>
            <Menu.Item key="4">
              <Link to="/contacts">Contacts</Link>
            </Menu.Item>
          </Menu>
        </div>

        <div>
          <ThemeSwitcher />
        </div>
      </Header>
    </Layout>
  );
};

export default TopMenu;
