import { Outlet } from "react-router-dom";

import { Navbar } from "@/features";

import "./Home.style.scss"

function Home() {

	console.log(`${Home.name} re-render`);
	return (
		<div id="home">
			<Navbar />
			<main>
				<Outlet />
			</main>
		</div>
	);
}

export default Home;
