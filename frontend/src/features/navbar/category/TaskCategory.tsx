import { useNavigate } from 'react-router-dom'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import Category from "./Category";
import tasks from "./tasks";

function TaskCategory() {
	const navigate = useNavigate();

	return (
		<Category
			data={tasks}
			header={"Tasks"}
			render={(task) => (
				<div 
					key={task.id} 
					className="categoryItem" 
					tabIndex={0}
					onClick={(event) => {
						event.preventDefault();
						navigate(`?taskType=${task.title}`)
					}}
				>
					<span>
						<FontAwesomeIcon icon={task.icon} className='icon'/>
					</span>
					<p className="categoryItemTitle">{task.title}</p>
				</div>
			)}
		/>
	);
}

export default TaskCategory;
