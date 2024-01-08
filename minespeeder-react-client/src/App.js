import { Routes, Route } from 'react-router-dom';
import Game from './pages/Game';
import Home from './pages/Home';

const App = () => {
 return (
    <>
       <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/game/:gameId" element={<Game />} />
       </Routes>
    </>
 );
};

export default App;