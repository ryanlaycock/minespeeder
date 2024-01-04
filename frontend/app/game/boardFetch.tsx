import useSWR from 'swr'

interface BoardResp {
    tiles: Tile[];
}

interface Tile {
  xPos: number;
  yPos: number;
  state: string;
}
let url = 'http://localhost:8080';
const fetcher = (url: string) => fetch(url).then(r => r.json())

export default function getBoard (gameId: string, boardId: string) {
    const { data, error, isLoading } = useSWR(
        `http://localhost:8080/v1/games/${gameId}/boards/${boardId}`,
        fetcher,
        { refreshInterval: 1000 }
    )
    if (error) {
        console.log(error);
    } else {
        console.log(data);
    }

    const board = data as BoardResp;
    return {
        boardResp: board,
        isLoading,
        isError: error
    }
}

// const Fetch = () => {
//   const [tiles, setTiles] = useState([]);
//   useEffect(() => {
//     fetch('http://localhost:8080/v1/games/game1/boards/board1')
//       .then((res) => {
//         return res.json();
//       })
//       .then((data) => {
//         console.log(data);
//         setTiles(data);
//       });
//   }, []);

  

//   return (
//     <div>
//         {/* {tiles.map((tile) => (
//         <div key={tile.id} style={{width: 20, height: 20, backgroundColor: 'gray', display: 'inline-block'}}>
//             {tile.x}, {tile.y}, {tile.state}
//         </div>
//         ))} */}


// {/*       
//       {photos.map((photo) => (
//         <img key={photo.id} src={photo.url} alt={photo.title} width={100} />
//       ))} */}
//     </div>
//   );
// };
// export default Fetch;