export type FlatListItem = {
	id: string;
	name: string;
};

export type FlatListProps<T> = {
	header?: React.ReactNode;
	data: T[];
	renderItem: (props: T) => React.ReactNode;
    className?: string
};

function FlatList<T extends FlatListItem>({
	header,
	data,
	renderItem,
    className = ""
}: FlatListProps<T>) {

	return (
		<div className={className}>
			<div>
                {header}
            </div>
			<div>{data.map((item) => renderItem(item))}</div>
		</div>
	);
}

export default FlatList;
