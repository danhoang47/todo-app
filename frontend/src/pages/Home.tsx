import { Outlet } from "react-router-dom";

import { MenuBar } from "@/features";

function Home() {
	console.log(`${Home.name} re-render`);
	return (
		<div className="flex">
			<MenuBar />
			<div className="xl:w-3/4 md:w-2/3">
				<Outlet />
			</div>
		</div>
	);
}

export default Home;
