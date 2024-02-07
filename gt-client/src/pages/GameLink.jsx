import { useParams } from "react-router-dom";

import { getGameLink } from "../api";

const GameLink = () => {
  let { id } = useParams();
  console.log(id);

  let links = [
    {
      ID: "f0000000-0000-0000-0000-000000000000",
      Name: "Tips & Tricks",
      Desc: "Link desc",
      URL: "http://example.com/link0",
      GameID: "00000000-0000-0000-0000-000000000000",
      CreatedAt: "2024-01-28T10:50:12.887143684+01:00",
      UpdatedAt: "2024-01-28T10:50:12.887143684+01:00",
      DeletedAt: null,
    },
    {
      ID: "f0000000-0000-0000-0000-000000000001",
      Name: "Tutorial for game 2",
      Desc: "Some cool link",
      URL: "http://example.com/link1",
      GameID: "00000000-0000-0000-0000-000000000000",
      CreatedAt: "2024-01-28T10:50:12.890323906+01:00",
      UpdatedAt: "2024-01-28T10:50:12.890323906+01:00",
      DeletedAt: null,
    },
  ];

  return (
    <>
      <h1>Game Link</h1>
      <div>
        {links.map((link) => (
          <div key={link.ID}>
            <p>{link.Desc}</p>
            <a href={link.URL} target="_blank" rel="noreferrer">
              {link.Name}
            </a>
          </div>
        ))}
      </div>
    </>
  );
};

export default GameLink;
