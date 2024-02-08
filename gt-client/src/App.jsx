import "./App.css";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Layout from "./Layout";
import Home from "./pages/Home";
import Games from "./pages/Games";
import GameLink from "./pages/GameLink";
import Login from "./pages/Login";

const router = createBrowserRouter([
  {
    path: "/",
    Component: Layout,
    children: [
      {
        path: "/",
        Component: Home,
      },
      {
        path: "games",
        Component: Games,
        children: [
          {
            path: "links/:id",
            Component: GameLink,
          },
        ],
      },
      {
        path: "register",
        Component: () => <div>Register</div>,
      },
      {
        path: "login",
        Component: Login,
      },
    ],
  },
]);

function App() {
  return <RouterProvider router={router} />;
}
export default App;
