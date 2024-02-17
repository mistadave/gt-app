import "./Games.css";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Link, Outlet } from "react-router-dom";
import {getGames} from "../api"
import { useEffect, useState } from "react";

const Games = () => {
  const [games, setGames] = useState([]);
 
  useEffect(() => {
    getGames().then((res) => {
      setGames(res);
    });
  }, []);
 
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
              <Badge variant="default">{game.Genre}</Badge>
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
