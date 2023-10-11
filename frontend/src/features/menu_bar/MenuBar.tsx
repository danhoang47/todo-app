import { memo } from "react";

function MenuBar() {
    

	console.log(`${MenuBar.name} re-render`);
	return <div className="xl:w-1/4 md:w-1/3">MenuBar</div>;
}

const MemoriedMenuBar = memo(MenuBar);
export default MemoriedMenuBar;
