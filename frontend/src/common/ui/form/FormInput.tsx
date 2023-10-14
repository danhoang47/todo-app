import { useState } from "react";

export type FormInputProps = {
	initialValue?: string | number;
	name: string;
	type?: React.HTMLInputTypeAttribute;
    className?: string;
};

function FormInput({ initialValue = "", name, type = "text", className = '' }: FormInputProps) {
	const [value, setValue] = useState(initialValue);

	const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		setValue(event.target.value);
	};

    const handleBlur = (event: React.FocusEvent<HTMLInputElement>) => {
        
    }

    const handleFocus = (event: React.FocusEvent<HTMLInputElement>) => {

    }

	return (
		<input 
            className={className}
            name={name} 
            type={type} 
            value={value} 
            onChange={handleChange} 
            onBlur={handleBlur}
            onFocus={handleFocus}
        />
	);
}

export default FormInput;
