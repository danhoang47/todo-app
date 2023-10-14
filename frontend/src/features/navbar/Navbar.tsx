import { memo, useState } from "react";
import { faAnglesLeft, faAnglesRight } from "@fortawesome/free-solid-svg-icons";

import { IconButton } from "@/common/ui";
import { TaskCategory, ListCategory } from "./category";
import NavbarHeader from "./header";
import NavbarFooter from "./footer";

import "./Navbar.styles.scss";

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
			<NavbarFooter />
		</nav>
	);
}

const MemoriedNavbar = memo(Navbar);
export default MemoriedNavbar;
