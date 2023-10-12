import { Outlet } from "react-router-dom";

import { Menu } from "@/features";

function Home() {
	console.log(`${Home.name} re-render`);
	return (
		<div className="flex">
			<Menu />
			<div className="flex-grow">
				<Outlet />
			</div>
		</div>
	);
}

export default Home;
