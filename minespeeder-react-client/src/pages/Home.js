import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';

export default function Home() {
    const navigate = useNavigate();
    const [numberOfBoards, setNumberOfBoards] = useState(2);
    const [height, setHeight] = useState(8);
    const [width, setWidth] = useState(8);
    const [numberOfBombs, setNumberOfBombs] = useState(5);

    function createNewGame(event) {
        event.preventDefault();
        const id = uuidv4();
        axios.post('http://localhost:8080/v1/games', { 
            "numberOfBoards": parseInt(numberOfBoards),
            "boardOptions": {
                "height": parseInt(height),
                "width": parseInt(width),
                "numberOfBombs": parseInt(numberOfBombs)
            },
            "id": id
        })
        .then(response => {
            navigate(`/game/${id}`);
        })
        .catch(error => {
            console.error('There was an error creating new game', error);
        });
    }

    return (
        <div className="home">
        <h1>💣 MineSpeeder</h1>
        <p>
            Minespeeder is a Minesweeper clone with a twist: it's multiplayer. 
            The first player to clear their board wins.
        </p>
        <p>
            Minespeeder is a work in progress. It's currently in a very early alpha state. 
            You can play it, but it's not very fun yet. 
            If you want to try it out, fill the form below to create a game.
        </p>
        <form onSubmit={createNewGame}>
            <label>
                Number of Boards:
                <input type="number" value={numberOfBoards} onChange={e => setNumberOfBoards(e.target.value)} />
            </label>
            <label>
                Height:
                <input type="number" value={height} onChange={e => setHeight(e.target.value)} />
            </label>
            <label>
                Width:
                <input type="number" value={width} onChange={e => setWidth(e.target.value)} />
            </label>
            <label>
                Number of Bombs:
                <input type="number" value={numberOfBombs} onChange={e => setNumberOfBombs(e.target.value)} />
            </label>
            <input type="submit" value="Create New Game" />
        </form>
        </div>
    );
}