import { Collapse } from "@/common/ui";
import { memo } from "react";

function MenuBar() {
    
	console.log(`${MenuBar.name} re-render`);
	return (
		<Collapse 
			visible={true} 
			className="xl:w-1/4 md:w-1/3 border-r dark:border-r-white light:border-r-light"
		>
			MenuBar
		</Collapse>
	);
}

const MemoriedMenuBar = memo(MenuBar);
export default MemoriedMenuBar;
