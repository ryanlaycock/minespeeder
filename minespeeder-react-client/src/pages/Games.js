import { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import Board from '../components/Board';

// Refactor to typescript

export default function Game() {
  const { gameId } = useParams();
  const [boards, setBoards] = useState();

  useEffect(() => {
    const interval = setInterval(() => {
      axios.get('http://localhost:8080/v1/games/' + gameId)
        .then(response => {
          setBoards(response.data.boards);
        })
        .catch(error => {
          console.error('There was an error loading latest board update', error);
        });
    }, 100);

    return () => clearInterval(interval);
  }, [gameId]);

  if (!boards) {
    return <div>Loading game...</div>;
  } 

  return (
    <div className="game">
      {boards.sort((a, b) => a.id.localeCompare(b.id)).map((board) => ( // hacky sort
        <div className="game-board">
          <Board board={board} key={board.id} id={board.id} gameId={gameId} />
        </div>
      ))}      
    </div>
  );
}
