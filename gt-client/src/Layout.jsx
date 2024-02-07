import React from "react";
import { Link, Outlet } from "react-router-dom";

const Layout = () => {
  return (
    <>
      <div>
        <h1>Games Tutor</h1>
        <p>
          Just some games
        </p>
        <div>
          <nav>
            <ul>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/games">Games</Link>
              </li>
            </ul>
          </nav>
          <hr />
        </div>
        <Outlet />
      </div>
    </>
  );
};

export default Layout;
