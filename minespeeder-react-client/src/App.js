import { useState, useEffect } from 'react';
import axios from 'axios';

function Tile({ value, onLeftClick, onRightClick }) { 
  return (
    <button 
      className="tile"
      onClick={onLeftClick}
      onContextMenu={(e) => {
        e.preventDefault(); // Prevent the context menu from showing up
        onRightClick();
      }}
    >
      {value}
    </button>
  );
}

function BoardRow({tilesRow, y, doAction}) {

  return (
    <div className="board-row">
      {tilesRow.map((tile, x) => (
        <Tile 
          key={x+","+y}
          value={tile}
          onLeftClick={() => doAction(y, x, actionType.REVEAL)}
          onRightClick={() => doAction(y, x, actionType.FLAG)}
        />
      ))}
    </div>
  )
}

let actionType = {
  REVEAL: "reveal",
  FLAG: "flag"
}

function Board({ width, height }) {
  var [tiles, setTiles] = useState(Array(width).fill(Array(height).fill("")));

  function doAction(x, y, actionType) {
    console.log(actionType + " " + x + " " + y);

    axios.post('http://localhost:8080/v1/games/game1/boards/board1/actions', {
      "xPos": x,
      "yPos": y,
      "type": actionType
    })
    .then(response => {
      console.log(response);
    });
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
            let value = "";
            switch (tile.state) {
              case "hidden":
                value = "❓";
                break;
              case "flag":
                value = "🚩";
                break;
              case "bomb":
                value = "💣";
                break;
              case "1":
                value = "1️⃣";
                break;
              case "2": 
                value = "2️⃣";
                break;
              case "3":
                value = "3️⃣";
                break;
              case "4":
                value = "4️⃣";
                break;
              case "5":
                value = "5️⃣"
                break;
              case "6":
                value = "6️⃣";
                break;
              case "7":
                value = "7️⃣";
                break;
              case "8":
                value = "8️⃣";
                break;
            }
            
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
        <BoardRow key={y} tilesRow={tile} y={y} doAction={doAction} />
      ))}
    </>
  );
}

export default function Game() {
  const [boardDimensions, setBoardDimensions] = useState(null);
  
  useEffect(() => {
    axios.get('http://localhost:8080/v1/games/game1/boards/board1')
      .then(response => {
        console.log(response);
        const height = response.data["height"];
        const width = response.data["width"];
        setBoardDimensions({ height, width });
        console.log("Board dimensions: " + height + " " + width);
      });
  }, []);

  if (!boardDimensions) {
    return <div>Loading game...</div>;
  }

  return (
    <div className="game">
      <div className="game-board">
        <Board height={boardDimensions.height} width={boardDimensions.width} />
      </div>
    </div>
  );
}