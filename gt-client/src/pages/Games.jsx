import "./Games.css";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Link, Outlet } from "react-router-dom";

const Games = () => {
  const games = [
    {
      ID: "f0000000-0000-0000-0000-000000000000",
      Name: "Game 0",
      Desc: "",
      Image: "",
      Genre: "",
      ReleaseDate: "2012-01-01T00:00:00Z",
      CreatedAt: "2024-01-28T10:50:12.917406203+01:00",
      UpdatedAt: "2024-01-28T10:50:12.917406203+01:00",
      DeletedAt: null,
    },
    {
      ID: "f0000000-0000-0000-0000-000000000001",
      Name: "Game 1",
      Desc: "",
      Image: "",
      Genre: "",
      ReleaseDate: "2014-02-21T00:00:00Z",
      CreatedAt: "2024-01-28T10:50:12.920291393+01:00",
      UpdatedAt: "2024-01-28T10:50:12.920291393+01:00",
      DeletedAt: null,
    },
  ];
  return (
    <>
      <div>
        <h1>Games</h1>
        <div className="grid">
          {games.map((game) => (
            <Card key={game.ID}>
              <CardHeader>
                <CardTitle>{game.Name}</CardTitle>
              </CardHeader>
              <CardContent>
                <img src={game.Image || "default-image.jpg"} alt="Game Image" />
                <Link to={`/games/links/${game.ID}`}>
                  <Button>{game.Name}</Button>
                </Link>
                <span>{game.ReleaseDate}</span>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
      <Outlet />
    </>
  );
};

export default Games;
