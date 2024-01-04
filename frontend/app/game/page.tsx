'use client';
import getBoard from "./boardFetch";

export default function Page() {
	return (
		<div>
            <p>💣 MineSpeeder 🏳️</p>
			<Board />
		</div>
	);
}

function Board () {
    const { boardResp, isLoading, isError } = getBoard("game1", "board1")
    if (isLoading) return <div>Loading...</div>
    if (isError) return <div>Error...</div>
    
    const rows = [];
    for (let y = 0; y < 4; y++) {
        const columns = [];
        for (let x = 0; x < 4; x++) {
            columns.push(Tile(x+","+y, "H"));
        }
        rows.push(<tr>{columns}</tr>);
    }
    return <table><tbody>{rows}</tbody></table>;
}

function Tile (id: string, state: string) {
    return (
        <td>
        <div key={id}>
            {state}
        </div>
        </td>
    )
}
