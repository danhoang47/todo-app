import { FontAwesomeIcon, FontAwesomeIconProps } from "@fortawesome/react-fontawesome";

export type IconButtonProps = FontAwesomeIconProps & {
    onClick?: (e?: React.MouseEvent) => void;
    buttonClassName?: string,
    children?: React.ReactNode
}

function IconButton({ onClick, buttonClassName = '', children, ...props }: IconButtonProps) {
    return (  
        <button onClick={onClick} className={buttonClassName}>
            <FontAwesomeIcon {...props}/>
            {children}
        </button>
    );
}

export default IconButton;