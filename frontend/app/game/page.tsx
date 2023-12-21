'use client';

export default function Page() {
    const callAPI = async () => {
        try {
            const res = await fetch(`http://localhost:8080/v1/games/game1/boards/board1`);
            const data = await res.json();
            console.log(data);
        } catch (err) {
            console.log(err);
        }
    };

	return (
		<div>
            <p>💣 MineSpeeder 🏳️</p>
			<main>
				<button onClick={callAPI}>Make API call</button>
			</main>
		</div>
	);
}

