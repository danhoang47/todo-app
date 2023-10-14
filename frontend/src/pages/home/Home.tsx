import { createPortal } from 'react-dom'
import { Outlet } from "react-router-dom";

import { Navbar, UserLoginSignup } from "@/features";

import "./Home.style.scss"

function Home() {

	console.log(`${Home.name} re-render`);
	return (
		<div id="home">
			<Navbar />
			<main>
				<Outlet />
			</main>
			{createPortal(<UserLoginSignup />, document.body)}
		</div>
	);
}

export default Home;
