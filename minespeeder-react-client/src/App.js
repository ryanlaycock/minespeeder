import { Routes, Route } from 'react-router-dom';
import Game from './pages/Games';

const App = () => {
 return (
    <>
       <Routes>
          <Route path="/game/:gameId" element={<Game />} />
       </Routes>
    </>
 );
};

export default App;