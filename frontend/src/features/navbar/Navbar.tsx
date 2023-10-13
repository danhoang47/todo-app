import { memo, useState } from "react";
import { faAnglesLeft, faAnglesRight } from "@fortawesome/free-solid-svg-icons";

import { IconButton } from "@/common/ui";
import TaskCategory from "./TaskCategory";
import ListCategory from "./ListCategory";

import "./Navbar.styles.scss";
import NavbarHeader from "./NavbarHeader";

function Navbar() {
	const [isNavbarExpand, setNavbarExpand] = useState(true);

	console.log(`${Navbar.name} re-render`);
	return (
		<nav
			id="navbar"
			className={`${isNavbarExpand ? "expand" : "collapse"}`}
		>
			<NavbarHeader />
			<TaskCategory />
			<ListCategory />
			<IconButton
				icon={isNavbarExpand ? faAnglesLeft : faAnglesRight}
				buttonClassName="expandMenuBtn"
				onClick={() => setNavbarExpand(!isNavbarExpand)}
			/>
		</nav>
	);
}

const MemoriedNavbar = memo(Navbar);
export default MemoriedNavbar;
