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
  
function BoardProgress({ boardProgress }) {
  return (
    <div className="board-progress">
      {boardProgress.state === "completed" && ( // TODO Understand this syntax :D 
        <div>You won!</div>
      )}
      {boardProgress.state === "failed" && ( // TODO Understand this syntax :D 
        <div>Exploded 💥 Start again!</div>
      )}
      <div className="board-progress-tiles">
        <div className="board-progress-tiles-remaining">
          {boardProgress.numOfRemainingTiles} / {boardProgress.numOfTiles} tiles remaining
        </div>
      </div>
      <div className="board-progress-bombs">
        <div className="board-progress-bombs-remaining">
          {boardProgress.numOfRemainingBombs} / {boardProgress.numOfBombs} bombs remaining
        </div>
      </div>
    </div>
  )
}
  
let actionType = {
  REVEAL: "reveal",
  FLAG: "flag"
}
  
export default function Board({ board, id, gameId }) {
  function doAction(x, y, actionType) {
    axios.post('http://localhost:8080/v1/games/' + gameId + '/boards/' + id + '/actions', {
      "xPos": x,
      "yPos": y,
      "type": actionType
    })
    .then(response => {
      console.log(response);
    })
    .catch(error => {
      console.error('There was an error sending actions', error);
    });;
  }

  let tilesCopy = Array(board.width).fill(null).map(() => Array(board.height).fill(null));
  for (let tile of board.tiles) {
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
      default:
        value = "";
        break;
    }
    tilesCopy[x][y] = value;
  }
  
  const boardProgress = {
    numOfTiles: board.numberOfTiles,
    numOfBombs: board.numberOfBombs,
    numOfRemainingTiles: board.numberOfRemainingTiles,
    numOfRemainingBombs: board.numberOfRemainingBombs,
    state: board.state
  }

  return (
    <>
      {tilesCopy.map((tile, y) => (
          <BoardRow key={y} tilesRow={tile} y={y} doAction={doAction} />
      ))}
      <BoardProgress boardProgress={boardProgress} />
    </>
  );
}
