import { useState, useEffect } from 'react';
import axios from 'axios';

function Tile({ value, onTileClick }) { 
  return (
    <button 
      className="tile"
      onClick={onTileClick}
    >
      {value}
    </button>
  );
}

function BoardRow({tilesRow, y, tileClick}) {

  return (
    <div className="board-row">
      {tilesRow.map((tile, x) => (
        <Tile key={x+","+y} value={tile} onTileClick={() => tileClick(y, x)} />
      ))}
    </div>
  )
}

export default function Board() {
  const [tiles, setTiles] = useState(Array(4).fill(Array(4).fill("🚩")));

  function setTile(x, y) {
    console.log("Clicked " + x + " " + y);
    const tilesCopy = tiles.map(row => [...row]);
    tilesCopy[x][y] = "💣";
    setTiles(tilesCopy);
  }

  useEffect(() => {
    const interval = setInterval(() => {
      axios.get('http://localhost:8080/v1/games/game1/boards/board1')
        .then(response => {
          const data = response.data["tiles"];
          
          const tilesCopy = tiles.map(row => [...row]);
          for (let tile of data) {
            let x = tile.xPos;
            let y = tile.yPos;
            let value = tile.state === "hidden" ? "🚩" : "💣";
            if (tilesCopy[x] != null && tilesCopy[x][y] != null) {
              tilesCopy[x][y] = value;
            }
          }

          setTiles(tilesCopy);
        });
    }, 100);

    return () => clearInterval(interval);
  }, []);
  
  return (
    <>
      {tiles.map((tile, y) => (
        <BoardRow key={y} tilesRow={tile} y={y} tileClick={setTile} />
      ))}
    </>
  );
}
