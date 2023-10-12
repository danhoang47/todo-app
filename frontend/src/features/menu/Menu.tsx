import { Collapse } from "@/common/ui";
import { memo } from "react";

function Menu() {
	console.log(`${Menu.name} re-render`);
	return (
		<Collapse 
			visible={true} 
			direction="horizontal" 
			className="md:w-1/4"
			invisibleClsx="md:w-20"
		>
			<div className="p-4 flex justify-between @container">
				{/* userinformation */}
				<p className="@xs:hidden">username</p>
				<Collapse.Trigger
					renderTrigger={(_, onClick) => (
						<button onClick={onClick}>Collapse</button>
					)}
				/>
			</div>
			<div>

			</div>
		</Collapse>
	);
}

const MemoriedMenu = memo(Menu);
export default MemoriedMenu;
