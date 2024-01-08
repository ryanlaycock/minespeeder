import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';
import { useNavigate } from 'react-router-dom';

export default function Home() {
    const navigate = useNavigate();
    
    function CreateNewGame() {
        const id = uuidv4();
        axios.post('http://localhost:8080/v1/games', { 
            "numberOfBoards": 2,
            "boardOptions": {
                "height": 8,
                "width": 8,
                "numberOfBombs": 5
            },
            "id": id
        })
        .then(response => {
            navigate(`/game/${id}`);
        })
        .catch(error => {
            console.error('There was an error creating new game', error);
        });;
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
            If you want to try it out, click the button below to create a game.
        </p>
        <button onClick={CreateNewGame} >New game</button>
        </div>
    );
}