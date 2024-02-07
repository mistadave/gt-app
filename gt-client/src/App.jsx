import "./App.css";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Layout from "./Layout";
import Home from "./pages/Home";
import Games from "./pages/Games";
import GameLink from "./pages/GameLink";

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
    ],
  },
]);

function App() {
  // const [count, setCount] = useState(0)

  return <RouterProvider router={router} />;
}
export default App;
