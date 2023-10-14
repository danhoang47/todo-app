import FormInput from "./FormInput";

export type FormProps<T> = {
	initialValues: T;
	errors?: Record<keyof T, string>;
	validate?: (values: T) => boolean;
	onSubmit: (value: T) => void;
    onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
    onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void;
    onFocus?: (event: React.FocusEvent<HTMLInputElement>) => void;
	handleSubmit?: (event: React.FormEvent<HTMLFormElement>) => void;
	handleChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
	handleBlur?: (event: React.FocusEvent<HTMLInputElement>) => void;
	handleFocus?: (event: React.FocusEvent<HTMLInputElement>) => void;
	children: (field: Omit<FormProps<T>, "children">) => React.ReactNode;
};

function Form<T extends Record<string, string | number>>({
	children,
	...props
}: FormProps<T>) {
	props.handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault();
		const formData = new FormData(event.currentTarget);
		const fields = Object.fromEntries(formData.entries())

		if (props.validate && props.validate(fields as T)) {
			props.onSubmit(fields as T);
		}

	};

    props.handleBlur = (event: React.FocusEvent<HTMLInputElement>) => {
        if (props.onBlur) props.onBlur(event)
    }

    props.handleFocus = (event: React.FocusEvent<HTMLInputElement>) => {
        if (props.onFocus) props.onFocus(event)
    }

	console.log(`Form re-render`)
	return children(props);
}

Form.Input = FormInput

export default Form;
