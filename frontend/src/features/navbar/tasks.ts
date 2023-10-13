import { faCalendarDays, faPaperPlane, faNoteSticky } from "@fortawesome/free-regular-svg-icons";
import { faListCheck } from "@fortawesome/free-solid-svg-icons";

const tasks = [
	{
		id: 1,
		icon: faPaperPlane,
		title: "upcoming",
	},
	{
		id: 2,
		icon: faListCheck,
		title: "today",
	},
	{
		id: 3,
		icon: faCalendarDays,
		title: "calendar",
	},
	{
		id: 4,
		icon: faNoteSticky,
		title: "sticky wall",
	},
];

export default tasks;