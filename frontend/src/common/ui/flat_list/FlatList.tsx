export type FlatListItem = {
	id: string;
	name: string;
};

export type FlatListProps<T> = {
	header?: React.ReactNode;
	data: T[];
	renderItem: (props: T) => React.ReactNode;
	collapsible?: boolean;
    className?: string
};

function FlatList<T extends FlatListItem>({
	header,
	data,
	renderItem,
	collapsible = false,
    className = ""
}: FlatListProps<T>) {

    const Wrapper = collapsible ? 'div' : 'a'

	return (
		<Wrapper className={className}>
			<div>
                {header}
            </div>
			<div>{data.map((item) => renderItem(item))}</div>
		</Wrapper>
	);
}

export default FlatList;
