import {
	useState,
	createContext,
	useCallback,
	useContext,
	useMemo,
} from "react";

export type CollapseProps = {
	visible: boolean;
	direction?: "vertical" | "horizontal";
	children: React.ReactNode;
	className?: string;
	visibleClsx?: string;
	invisibleClsx?: string;
};

export type CollapseContextType = {
	isVisible: boolean;
	onTriggerButtonClick: () => void;
};

const CollapseContext = createContext<CollapseContextType | undefined>(
	undefined
);

function CollapseTrigger({
	renderTrigger,
}: {
	renderTrigger: (isVisible: boolean, onClick: () => void) => React.ReactNode;
}) {
	const { isVisible, onTriggerButtonClick } = useContext(CollapseContext)!;

	return <>{renderTrigger(isVisible, onTriggerButtonClick)}</>;
}

function Collapse({
	visible,
	children,
	className = "",
	visibleClsx = "",
	invisibleClsx = "",
}: CollapseProps) {
	const [isVisible, setVisible] = useState(visible);
	const expandClassName = useMemo(() => isVisible ? visibleClsx : invisibleClsx, [isVisible])

	const onTriggerButtonClick = useCallback(() => {
		setVisible((prev) => !prev);
	}, []);

	const collapseContextValue = useMemo(
		() => ({
			isVisible,
			onTriggerButtonClick,
		}),
		[isVisible, onTriggerButtonClick]
	);

	return (
		<CollapseContext.Provider value={collapseContextValue}>
			<div
				className={`transition-all overflow-x-hidden ${className} ${expandClassName}`}
			>
				{children}
			</div>
		</CollapseContext.Provider>
	);
}

Collapse.Trigger = CollapseTrigger;

export default Collapse;
