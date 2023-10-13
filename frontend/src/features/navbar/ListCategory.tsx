import Category from "./Category";

const lists = [
	{
		id: 1,
		color: "#ff6b6b",
		title: "personal",
	},
	{
		id: 2,
		color: "#66d9e8",
		title: "Work",
	},
	{
		id: 3,
		color: "#ffd43b",
		title: "List 1",
	},
];

function ListCategory() {
	return (
		<Category
			header={"Lists"}
			data={lists}
			render={(list) => (
				<div key={list.id} className="categoryItem" tabIndex={0}>
					<span
						className="listItemColor"
						style={{ backgroundColor: list.color }}
					/>
					<p className="categoryItemTitle">{list.title}</p>
				</div>
			)}
		/>
	);
}

export default ListCategory;
