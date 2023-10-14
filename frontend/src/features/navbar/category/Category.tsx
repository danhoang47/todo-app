
import "./Category.styles.scss"

export type CategoryProps<T> = {
    header?: React.ReactNode,
    data: T[],
    render: (item: T) => React.ReactNode,
	extraItem?: React.ReactNode,
	className?: string
};

function Category<T>({ header, data, render, extraItem, className = "" }: CategoryProps<T>) {
	return (
		<div className={`categoryOptions ${className}`.trim()}>
			<div className="categoryHeader">{header}</div>
			<div className="categoryItemList">
				{data.map((item) => render(item))}
				{extraItem}
			</div>
		</div>
	);
}

export default Category;
