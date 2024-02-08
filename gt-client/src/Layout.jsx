import { Outlet } from "react-router-dom";
import Header from "./pages/Header";

const Layout = () => {
  return (
    <>
      <nav>
        <Header></Header>
      </nav>
      <div>
        <h1>Games Tutor</h1>
        <hr></hr>
        <Outlet />
      </div>
    </>
  );
};

export default Layout;
